package services

import (
	"context"
	"errors"
	"ielts-app-api/common"
	"ielts-app-api/internal/models"
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

	// Hash the password
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

func (s *Service) LoginUser(ctx context.Context, req models.LoginRequest) (string, error) {
	user, err := s.UserRepo.GetDetailByConditions(ctx, func(tx *gorm.DB) {
		tx.Where("email = ?", req.Email)
	})
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.RoleID,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
