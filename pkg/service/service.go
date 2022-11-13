package service

import "github.com/Big0ak/delivery-company/models"

//import "github.com/Big0ak/delivery-company/pkg/repository"

// type Authorization interface {
// 	CreateUser(models.Manager) (int, error)
// }

// type Orders interface {
// }


type Service struct {
	*AuthService
	*OrderService
	repos Repository
}

////////////////////////////////////////////////////////////////////
type AuthorizationDB interface{
	CreateNewManagerDB(models.Manager) (int, error)
	GetManager(managerLogin, password string) (models.Manager, error)
	
	CreateNewClientDB(models.Client) (int, error)
	GetClient(clientLogin, password string) (models.Client, error)
}

type OrderDB interface{
	CreateManagerDB(managerId int, order models.Orders) (int, error)
	GetAllDB(managerId int) ([]models.Orders, error)
	GetByIdManagerDB(managerid, id int) (models.Orders, error)
	DeleteManagerDB(managerid, id int) error
	UpdateManagerDB(managerid, id int, input models.Orders) error 
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