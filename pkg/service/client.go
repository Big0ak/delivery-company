package service

import "github.com/Big0ak/delivery-company/models"

type ClientService struct {
	repo ClientDB
}

func NewClientService(repo ClientDB) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) GetAllClient(managerId int) ([]models.Client, error) {
	return s.repo.GetAllClientDB(managerId)
}