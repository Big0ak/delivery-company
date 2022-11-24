package service

import "github.com/Big0ak/delivery-company/models"

type CabinetService struct {
	repo CabinetDB
}

func NewCabinetService(repo CabinetDB) *CabinetService {
	return &CabinetService{repo: repo}
}

func (s *CabinetService) GetInfoManager(managerId int) (models.Manager, error){
	return s.repo.GetInfoManagerDB(managerId)
}

func (s *CabinetService) UpdateManager(managerId int, manager models.Manager) error{
	manager.Password = generatePasswordHash(manager.Password)
	return s.repo.UpdateManagerDB(managerId, manager)
}

func (s *CabinetService) GetInfoClient(clientId int) (models.Client, error) {
	return s.repo.GetInfoClientDB(clientId)
}

func (s *CabinetService) UpdateClient(clientId int, client models.Client) error{
	client.Password = generatePasswordHash(client.Password)
	return s.repo.UpdateClientDB(clientId, client)
}
