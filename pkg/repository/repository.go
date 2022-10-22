package repository

import (
	"database/sql"

	//"github.com/Big0ak/delivery-company/models"
)

// type AuthorizationDB interface{
// 	CreateUserDB(models.Manager) (int, error)
// }

// type OrdersDB interface{

// }

type RouteDB interface{

}

type Repository struct{
	*AuthSQLServer
	*OrderDB
	RouteDB
}

func NewRepository (db *sql.DB) *Repository{
	return &Repository{
		AuthSQLServer: NewAuthQLServer(db),
		OrderDB: NewOrderDB(db),
	}
}