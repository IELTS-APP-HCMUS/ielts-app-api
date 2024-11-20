package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	_, err := s.UserRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
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
		user, err := s.UserRepo.Create(ctx, &newUser)
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
		_, err = s.TargetRepo.Create(ctx, &newUserTarget)
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

		user, err = s.UserRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
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
				user, err = s.UserRepo.Create(ctx, &newUser)
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
				_, err = s.TargetRepo.Create(ctx, &newUserTarget)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
	} else {
		user, err = s.UserRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
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

	expiry, err := common.NormalizeToBangkokTimezone(time.Now().Add(5 * time.Minute))
	if err != nil {
		return "", err
	}

	newOTP := models.OTP{
		Email:  email,
		OTP:    otp,
		Expiry: expiry,
	}

	_, err = s.OTPRepo.Create(ctx, &newOTP)
	if err != nil {
		return "", err
	}

	return otp, nil
}

func (s *Service) ValidateOTP(ctx context.Context, email, otp string) error {
	// Fetch the stored OTP for the given email
	storedOTP, err := s.OTPRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("email = ?", email)
	})
	if err != nil {
		return errors.New("invalid OTP")
	}

	// Check if OTP has expired
	currentTime, err := common.NormalizeToBangkokTimezone(time.Now())
	if err != nil {
		return err
	}
	expiryTime, err := common.NormalizeToBangkokTimezone(storedOTP.Expiry)
	if err != nil {
		return err
	}

	fmt.Println("current time: ", currentTime)
	fmt.Println("expiry time: ", expiryTime)

	if expiryTime.Before(currentTime) {
		return errors.New("OTP has expired")
	}

	if storedOTP.OTP != otp {
		return errors.New("invalid OTP")
	}

	return nil
}

func (s *Service) ResetPassword(ctx context.Context, email, newPassword string) error {
	_, err := s.UserRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("email = ?", email)
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("email not found")
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

	return s.UserRepo.UpdatesByConditions(ctx, &updatedUser, func(tx *gorm.DB) {
		tx.Where("email = ?", email)
	})
}
