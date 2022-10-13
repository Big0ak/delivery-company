package service

//import "github.com/Big0ak/DeliveryCompany/pkg/repository"

type Authorization interface {
}

type Orders interface {
}

type Route interface {
}

type Service struct {
	Authorization
	Orders
	Route
	repo Repository
}

////////////////////////////////////////////////////////////////////
type Repository interface{

}

func NewService(repo Repository) *Service {
	return &Service{repo: repo} 
}

// func NewService(repos *repository.Repository) *Service {
// 	return &Service{}
// }
/////////////////////////////////////////////////////////////////////