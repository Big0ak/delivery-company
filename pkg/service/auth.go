package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Big0ak/delivery-company/models"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "uf3b289g38bf83nf3"
	signingKey = "g43g4nii#23523f3j3i2r"
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo AuthorizationDB
}

type tokenClaims struct {
	jwt.StandardClaims
	ManagerId int `json:"manager_id"`
}

func NewAuthServise(repo AuthorizationDB) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(manager models.Manager) (int, error) {
	manager.Password = s.generatePasswordHash(manager.Password)
	return s.repo.CreateUserDB(manager)
}

func (s *AuthService) GenerateToken(managerLogin, password string) (string, error) {
	user, err := s.repo.GetUser(managerLogin, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), // время жизни токена
			IssuedAt: time.Now().Unix(), // время создания токена
		},
		user.ManagerID,
	})

	return token.SignedString([]byte(signingKey))
}


func (s *AuthService) ParseToken(accessToken string) (int, error) {
	// на вход структуру Claims функцию возвращает ключ подписи или ошибку
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing mathod")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.ManagerId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}