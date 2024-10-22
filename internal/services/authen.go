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

	if req.AccessToken != nil {
		googleUser, err := verifyGoogleOAuthToken(*req.AccessToken)
		if err != nil {
			return nil, err
		}

		googleUserProfile, err := fetchGoogleUserProfile(*req.IdToken)
		if err != nil {
			return nil, err
		}
		fmt.Println("Flag 01: ", googleUserProfile)

		user, err = s.UserRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
			tx.Where("email = ? OR provider= ?", googleUserProfile.Email, common.USER_PROVIDER_GOOGLE)
		})
		fmt.Println("Flag 02: ")
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newUser := models.User{
					FirstName: &googleUserProfile.GivenName,
					LastName:  &googleUserProfile.FamilyName,
					Email:     googleUser.Email,
					RoleID:    common.ROLE_END_USER_UUID,
					Provider:  common.USER_PROVIDER_GOOGLE,
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
		fmt.Println("Flag 03: ")
		user, err = s.UserRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
			tx.Where("email = ?", req.Email)
		})
		fmt.Println("Flag 04: ")

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

	// Log the response body for debugging purposes
	fmt.Println("Response Body:", string(bodyBytes))

	// It returns 400 if the token is invalid
	if resp.StatusCode != http.StatusOK {
		return nil, common.ErrInvalidGoogleAuthenToken
	}

	// Decode the response body into the googleUser struct
	var googleUser models.GoogleUser
	if err := json.Unmarshal(bodyBytes, &googleUser); err != nil {
		fmt.Println("Error 23: ", err)
		return nil, err
	}

	return &googleUser, nil
}

func fetchGoogleUserProfile(accessToken string) (*models.GoogleUserProfile, error) {

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body into a variable
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Log the response body for debugging purposes
	fmt.Println("Response Body:", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch Google user profile")
	}

	var googleUserProfile models.GoogleUserProfile
	if err := json.Unmarshal(bodyBytes, &googleUserProfile); err != nil {
		return nil, err
	}

	return &googleUserProfile, nil
}
