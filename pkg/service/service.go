package service

import "github.com/Big0ak/delivery-company/models"

//import "github.com/Big0ak/delivery-company/pkg/repository"

// type Authorization interface {
// 	CreateUser(models.Manager) (int, error)
// }

type Orders interface {
}

type Route interface {
}

type Service struct {
	*AuthService
	Orders
	Route
	repos Repository
}

////////////////////////////////////////////////////////////////////
type AuthorizationDB interface{
	CreateUserDB(models.Manager) (int, error)
	GetUser(managerLogin, password string) (models.Manager, error)
}

type Repository interface{
	AuthorizationDB
}

func NewService(repos Repository, AuthDB AuthorizationDB) *Service {
	return &Service{
		repos: repos,
		AuthService: NewAuthServise(AuthDB)} 
}

// func NewService1(repos *repository.Repository) *Service {
// 	return &Service{
// 		Authorization: NewAuthServise(repos.AuthorizationDB),
// 	}
// }
/////////////////////////////////////////////////////////////////////