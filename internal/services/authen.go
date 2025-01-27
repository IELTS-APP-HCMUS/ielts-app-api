package services

import (
	"context"
	"encoding/json"
	"errors"
	"ielts-app-api/common"
	"ielts-app-api/internal/models"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JWTSecret = []byte("your_secret_key")

func (s *Service) SignupUser(ctx context.Context, req models.SignupRequest) error {

	_, err := s.userRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("email = ?", req.Email)
	})

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if req.Role == common.ROLE_END_USER {
		newUser := models.User{
			Email:     req.Email,
			Password:  string(hashedPassword),
			RoleID:    common.ROLE_END_USER_UUID,
			FirstName: &req.FirstName,
			LastName:  &req.LastName,
		}
		user, err := s.userRepo.Create(ctx, &newUser)
		if err != nil {
			return err
		}
		defaultDate := "1900-01-01" // default date
		parsedTime, err := time.Parse(time.DateOnly, defaultDate)
		if err != nil {
			return err
		}
		newUserTarget := models.Target{
			ID:                  user.ID,
			TargetStudyDuration: 0,
			TargetReading:       -1,
			TargetListening:     -1,
			TargetSpeaking:      -1,
			TargetWriting:       -1,
			NextExamDate:        parsedTime,
		}
		_, err = s.targetRepo.Create(ctx, &newUserTarget)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) LoginUser(ctx context.Context, req models.LoginRequest) (*string, error) {
	var user *models.User
	var err error

	if req.IdToken != nil {
		googleUser, err := verifyGoogleOAuthToken(*req.IdToken)
		if err != nil {
			return nil, err
		}

		user, err = s.userRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
			tx.Where("email = ? AND provider= ?", googleUser.Email, common.USER_PROVIDER_GOOGLE)
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newUser := models.User{
					FirstName: &googleUser.GivenName,
					LastName:  &googleUser.FamilyName,
					Email:     googleUser.Email,
					RoleID:    common.ROLE_END_USER_UUID,
					Provider:  common.USER_PROVIDER_GOOGLE,
					IsActive:  true,
				}
				user, err = s.userRepo.Create(ctx, &newUser)
				if err != nil {
					return nil, err
				}
				defaultDate := "1900-01-01" // default date
				parsedTime, err := time.Parse(time.DateOnly, defaultDate)
				if err != nil {
					return nil, err
				}
				newUserTarget := models.Target{
					ID:                  user.ID,
					TargetStudyDuration: 0,
					TargetReading:       -1,
					TargetListening:     -1,
					TargetSpeaking:      -1,
					TargetWriting:       -1,
					NextExamDate:        parsedTime,
				}
				_, err = s.targetRepo.Create(ctx, &newUserTarget)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
	} else {
		user, err = s.userRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
			tx.Where("email = ?", req.Email)
		})

		if err != nil {
			return nil, err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*req.Password)); err != nil {
			return nil, common.ErrInvalidEmailOrPassWord
		}
	}

	return generateJWTToken(user)
}

func generateJWTToken(user *models.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.RoleID,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func verifyGoogleOAuthToken(idToken string) (*models.GoogleUser, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=" + idToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, common.ErrInvalidGoogleAuthenToken
	}
	var googleUser models.GoogleUser
	if err := json.Unmarshal(bodyBytes, &googleUser); err != nil {
		return nil, err
	}
	return &googleUser, nil
}

func (s *Service) GenerateOTP(ctx context.Context, email string) (string, error) {
	otp := common.GenerateRandomOTP()

	expiry := time.Now().UTC().Add(1 * time.Minute)

	existingOTP, err := s.OTPRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("target = ? AND type = ?", email, common.TypeResetPassword)
	})

	if err == nil {
		existingOTP.IsVerified = true
		_, err = s.OTPRepo.Update(ctx, existingOTP.ID, existingOTP)
		if err != nil {
			return "", common.ErrFailedToInValidateExistingOTP
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	newOTP := models.OTP{
		Target:     email,
		Type:       common.TypeResetPassword,
		OTPCode:    otp,
		ExpiredAt:  expiry,
		IsVerified: false,
	}

	_, err = s.OTPRepo.Create(ctx, &newOTP)
	if err != nil {
		return "", err
	}

	return otp, nil
}

func (s *Service) ValidateOTP(ctx context.Context, email, otp string) error {
	storedOTP, err := s.OTPRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("target = ? AND type = ?", email, common.TypeResetPassword)
		tx.Order("created_at desc")
	})
	if err != nil {
		return err
	}
	if storedOTP.IsVerified {
		return common.ErrOTPAlreadyVerified
	}

	expiryTime, err := common.NormalizeToBangkokTimezone(storedOTP.ExpiredAt)
	if err != nil {
		return err
	}
	currentTime, err := common.NormalizeToBangkokTimezone(time.Now())
	if err != nil {
		return err
	}

	newAttempt := models.OTPAttempt{
		OTPID:     storedOTP.ID,
		Value:     otp,
		IsSuccess: false,
		CreatedAt: currentTime,
	}

	if expiryTime.Before(currentTime) {
		newAttempt.IsSuccess = false
		_, _ = s.OTPAttemptRepo.Create(ctx, &newAttempt)
		return common.ErrOTPExpired
	}

	if storedOTP.OTPCode != otp {
		newAttempt.IsSuccess = false
		_, _ = s.OTPAttemptRepo.Create(ctx, &newAttempt)
		return common.ErrInvalidOTP
	}

	storedOTP.IsVerified = true
	_, err = s.OTPRepo.Update(ctx, storedOTP.ID, storedOTP)
	if err != nil {
		return common.ErrFailedToUpdateOTPStatus
	}

	newAttempt.IsSuccess = true
	_, _ = s.OTPAttemptRepo.Create(ctx, &newAttempt)

	return nil
}

func (s *Service) ResetPassword(ctx context.Context, email, newPassword string) error {
	_, err := s.userRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("email = ?", email)
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.ErrEmailNotFound
		}
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	updatedUser := models.User{
		Password: string(hashedPassword),
	}

	return s.userRepo.UpdatesByConditions(ctx, &updatedUser, func(tx *gorm.DB) {
		tx.Where("email = ?", email)
	})
}
