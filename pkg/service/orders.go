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

func (s *OrderService) GetAllWithStatus(managerId int, status string) ([]models.OrdersRead, error) {
	return s.repo.GetAllWithStatusDB(managerId, status)
}

func (s *OrderService) GetById(userId, id int) (models.OrdersRead, error) {
	return s.repo.GetByIdDB(userId, id)
}

func (s *OrderService) DeleteManager(managerid, id int) error {
	return s.repo.DeleteManagerDB(managerid, id)
}

func (s *OrderService) UpdateOrderManager(managerid, id int, input models.Orders) error {
	return s.repo.UpdateOrderManagerDB(managerid, id, input)
}

// -------------------------------------------------------------------------------------------------
// ------------------------------ Client function --------------------------------------------------
// -------------------------------------------------------------------------------------------------


func (s *OrderService) GetUserOrder(clientId int) ([]models.OrdersRead, error) {
	return s.repo.GetUserOrderDB(clientId)
}

func (s *OrderService) GetAllWithStatusUserDB(clientId int, status string) ([]models.OrdersRead, error) {
	return s.repo.GetAllWithStatusUserDB(clientId, status)
}

func (s *OrderService) SearchOrdersByCityManager(managerId int, city string) ([]models.OrdersRead, error) {
	return s.repo.SearchOrdersByCityManagerDB(managerId, city)
}

func (s *OrderService) SearchOrdersByCityClient(clientId int, city string) ([]models.OrdersRead, error) {
	return s.repo.SearchOrdersByCityClientDB(clientId, city)
}
