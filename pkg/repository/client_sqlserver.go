package repository

import (
	"database/sql"
	"fmt"

	"github.com/Big0ak/delivery-company/models"
)

type ClientDB struct {
	db *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB{
	return &ClientDB{db: db}
}

func (r *ClientDB) GetAllClientDB(managerId int) ([]models.Client, error) {
	var client []models.Client
	query := fmt.Sprintf("SELECT c.ClientID, c.UserLogin, c.Name, c.Surname, c.Phone, c.RegistrationDate FROM %s c", clientTable)
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for row.Next(){
		c := models.Client{}
		err = row.Scan(&c.ClientID, &c.UserLogin, &c.Name, &c.Surname, &c.Phone, &c.RegistrationDate)
		if err != nil {
			return nil, err
		}
		client = append(client, c)
	}

	return client, nil
}