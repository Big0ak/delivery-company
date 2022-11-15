package service

import "github.com/Big0ak/delivery-company/models"

type DriverService struct {
	repo DriverDB
}

func NewDriverService(repo DriverDB) *DriverService {
	return &DriverService{repo: repo}
}

func (s *DriverService) GetAllDriver(managerId int) ([]models.Driver, error){
	return s.repo.GetAllDriverDB(managerId)
} 