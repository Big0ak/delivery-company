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
	signingKeyManager = "g43g4nii#23523f3j3i2r"
	signingKeyClient = "vu43v98be2io3bvv^be4"
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo AuthorizationDB
}

func NewAuthServise(repo AuthorizationDB) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateNewManager(manager models.Manager) (int, error) {
	manager.Password = s.generatePasswordHash(manager.Password)
	return s.repo.CreateNewManagerDB(manager)
}

func (s *AuthService) CreateNewClient(client models.Client) (int, error) {
	client.Password = s.generatePasswordHash(client.Password)
	return s.repo.CreateNewClientDB(client)
}

func (s *AuthService) GenerateTokenManager(managerLogin, password string) (string, error) {
	manager, err := s.repo.GetManager(managerLogin, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaimsManager{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), // время жизни токена
			IssuedAt: time.Now().Unix(), // время создания токена
		},
		manager.ManagerID,
	})

	return token.SignedString([]byte(signingKeyManager))
}

func (s *AuthService) GenerateTokenClient(clientLogin, password string) (string, error) {
	client, err := s.repo.GetClient(clientLogin, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaimsClient{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		client.ClientID,
	})

	return token.SignedString([]byte(signingKeyClient))
}

type tokenClaimsManager struct {
	jwt.StandardClaims
	ManagerId int `json:"manager_id"`
}

func (s *AuthService) ParseTokenManager(accessToken string) (int, error) {
	// на вход структуру Claims функцию возвращает ключ подписи или ошибку
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaimsManager{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing mathod")
		}

		return []byte(signingKeyManager), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaimsManager)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.ManagerId, nil
}

type tokenClaimsClient struct {
	jwt.StandardClaims
	ClientId int `json:"client_id"`
}

func (s *AuthService) ParseTokenClient(accessToken string) (int, error) {
	// на вход структуру Claims функцию возвращает ключ подписи или ошибку
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaimsClient{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing mathod")
		}

		return []byte(signingKeyClient), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaimsClient)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.ClientId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}