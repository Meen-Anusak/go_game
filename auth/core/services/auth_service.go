package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go_game/auth/core/domain"
	"go_game/auth/core/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type AuthService interface {
	Login(login domain.Login) (*domain.Auth, error)

	Register(user domain.User) error
}
type AuthServiceImpl struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) Login(login domain.Login) (*domain.Auth, error) {
	user, err := s.repo.Login(login)
	if err != nil {
		return nil, err
	}

	if ok := comparePasswords(user.Password, []byte(login.Password)); !ok {
		return nil, errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	var userResponse domain.Auth

	userResponse.UserName = user.UserName
	userResponse.Email = user.Email
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.Token = tokenString
	userResponse.Expired = time.Now().Add(5 * time.Minute)

	return &userResponse, nil
}

func (s *AuthServiceImpl) Register(user domain.User) error {
	password, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	if err := s.repo.Register(user); err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
