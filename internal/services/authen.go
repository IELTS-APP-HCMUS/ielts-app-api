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
		_, err = s.UserRepo.Create(ctx, &newUser)
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
			tx.Where("email = ? OR provider= ?", googleUser.Email, common.USER_PROVIDER_GOOGLE)
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
