package service

import (
	"crypto/sha1"
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
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.ManagerID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}