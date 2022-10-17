package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Big0ak/delivery-company/models"
)

const salt = "uf3b289g38bf83nf3"

type AuthService struct {
	repo AuthorizationDB
}

func NewAuthServise(repo AuthorizationDB) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(manager models.Manager) (int, error) {
	manager.Password = s.generatePasswordHash(manager.Password)
	return s.repo.CreateUserDB(manager)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}