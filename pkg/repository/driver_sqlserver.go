package repository

import (
	"database/sql"
	"fmt"

	"github.com/Big0ak/delivery-company/models"
)

type DriverDB struct {
	db *sql.DB
}

func NewDriverDB(db *sql.DB) *DriverDB {
	return &DriverDB{db: db}
}

func (d *DriverDB) GetAllDriverDB(managerId int) ([]models.Driver, error){
	var drivers []models.Driver
	query := fmt.Sprintf("SELECT * FROM %s", driverTable)
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		d := models.Driver{}
		err := rows.Scan(&d.DriverID, &d.Name, &d.Surname)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, d)
	}

	return drivers, nil
}