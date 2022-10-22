package service

import "github.com/Big0ak/delivery-company/models"

//import "github.com/Big0ak/delivery-company/pkg/repository"

// type Authorization interface {
// 	CreateUser(models.Manager) (int, error)
// }

// type Orders interface {
// }

type Route interface {
}

type Service struct {
	*AuthService
	*OrderService
	Route
	repos Repository
}

////////////////////////////////////////////////////////////////////
type AuthorizationDB interface{
	CreateManagerDB(models.Manager) (int, error)
	GetManager(managerLogin, password string) (models.Manager, error)
}

type OrderDB interface{
	CreateDB(managerId int, order models.Orders) (int, error)
	GetAllDB(managerId int) ([]models.Orders, error)
}

type Repository interface{
	AuthorizationDB
	OrderDB
}

func NewService(repos Repository, AuthDB AuthorizationDB, OrderDB OrderDB) *Service {
	return &Service{
		repos: repos,
		AuthService: NewAuthServise(AuthDB),
		OrderService: NewOrderService(OrderDB),
	} 
}

// func NewService1(repos *repository.Repository) *Service {
// 	return &Service{
// 		Authorization: NewAuthServise(repos.AuthorizationDB),
// 	}
// }
/////////////////////////////////////////////////////////////////////