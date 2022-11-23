package repository

import (
	"database/sql"
)

type Repository struct{
	*AuthSQLServer
	*OrderDB
	*ClientDB
	*DriverDB
	*CabinetDB
}

func NewRepository (db *sql.DB) *Repository{
	return &Repository{
		AuthSQLServer: NewAuthQLServer(db),
		OrderDB: NewOrderDB(db),
		ClientDB: NewClientDB(db),
		DriverDB: NewDriverDB(db),
		CabinetDB: NewCabinetDB(db),
	}
}