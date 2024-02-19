package service

import (
	"go_game/auth/core/domain"
	"go_game/auth/core/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	Login(login domain.Login) error

	Register(user domain.User) error
}
type AuthServiceImpl struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) Login(login domain.Login) error {
	panic("login")
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

//func GetToken() string,error {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
//		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
//	})
//
//	ss, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
//	if err != nil {
//		return nil,err
//	}
//
//	return ss,err
//
//}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
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
