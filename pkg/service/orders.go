package service

import "github.com/Big0ak/delivery-company/models"

type OrderService struct {
	repo OrderDB
}

func NewOrderService(repo OrderDB) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateManager(managerId int, order models.Orders) (int, error) {
	return s.repo.CreateManagerDB(managerId, order)
}

func (s *OrderService) GetAll(managerId int) ([]models.OrdersRead, error) {
	return s.repo.GetAllDB(managerId)
}

func (s *OrderService) GetByIdManager(managerid, id int) (models.OrdersRead, error) {
	return s.repo.GetByIdManagerDB(managerid, id)
}

func (s *OrderService) DeleteManager(managerid, id int) error {
	return s.repo.DeleteManagerDB(managerid, id)
}

func (s *OrderService) UpdateManager(managerid, id int, input models.Orders) error {
	return s.repo.UpdateManagerDB(managerid, id, input)
}
