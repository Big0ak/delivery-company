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