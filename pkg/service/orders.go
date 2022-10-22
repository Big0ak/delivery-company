package service

import "github.com/Big0ak/delivery-company/models"

type OrderService struct {
	repo OrderDB
}

func NewOrderService(repo OrderDB) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(managerId int, order models.Orders) (int, error) {
	return s.repo.CreateDB(managerId, order)
}

func (s *OrderService) GetAll(managerId int) ([]models.Orders, error) {
	return s.repo.GetAllDB(managerId)
}
