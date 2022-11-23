package service

import "github.com/Big0ak/delivery-company/models"

type CabinetService struct {
	repo CabinetDB
}

func NewCabinetService(repo CabinetDB) *CabinetService {
	return &CabinetService{repo: repo}
}

func (s *CabinetService) GetInfoClient(clientId int) (models.Client, error) {
	return s.repo.GetInfoClientDB(clientId)
}