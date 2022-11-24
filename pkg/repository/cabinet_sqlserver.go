package repository

import (
	"database/sql"
	"fmt"

	"github.com/Big0ak/delivery-company/models"
)

type CabinetDB struct {
	db *sql.DB
}

func NewCabinetDB(db *sql.DB) *CabinetDB{
	return &CabinetDB{db: db}
}

func (r *CabinetDB) GetInfoManagerDB(managerId int) (models.Manager, error){
	var manager models.Manager
	query := fmt.Sprintf("SELECT m.ManagerID, m.ManagerLogin, m.Name, m.Surname FROM %s m WHERE m.ManagerID = %d",
		managerTable, managerId)
	row := r.db.QueryRow(query)
	err := row.Scan(&manager.ManagerID, &manager.ManagerLogin, &manager.Name, &manager.Surname)
	if err != nil {
		return models.Manager{}, err
	}
	return manager, nil
} 

func (r *CabinetDB) UpdateManagerDB(managerId int, manager models.Manager) error{
	query := fmt.Sprintf("UPDATE %s SET Password = '%s', Name = '%s', Surname = '%s' WHERE ManagerID = %d",
		managerTable, manager.Password, manager.Name, manager.Surname, managerId)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *CabinetDB) GetInfoClientDB(clientId int) (models.Client, error){
	var client models.Client
	query := fmt.Sprintf("SELECT c.ClientID, c.UserLogin, c.Name, c.Surname, c.Phone, c.RegistrationDate FROM %s c WHERE c.ClientID = %d", 
		clientTable, clientId)
	row := r.db.QueryRow(query)
	err := row.Scan(&client.ClientID, &client.UserLogin, &client.Name, &client.Surname, &client.Phone, &client.RegistrationDate)
	if err != nil{
		return models.Client{}, err
	}
	return client, nil
}

func (r *CabinetDB) UpdateClientDB(clientId int, client models.Client) error{
	query := fmt.Sprintf("UPDATE %s SET Password = '%s', Name = '%s', Surname = '%s', Phone = '%s' WHERE ClientID = %d",
		clientTable, client.Password, client.Name, client.Surname, client.Phone, clientId)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}