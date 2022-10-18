package repository

import (
	"database/sql"
	"fmt"

	"github.com/Big0ak/delivery-company/models"
)

type AuthSQLServer struct {
	db *sql.DB
}

func NewAuthQLServer(db *sql.DB) *AuthSQLServer {
	return &AuthSQLServer{db: db}
}

func (r *AuthSQLServer) CreateUserDB(manager models.Manager) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (ManagerLogin, Password, Name, Surname) OUTPUT Inserted.ManagerID VALUES ('%s', '%s', '%s', '%s')", managerTable, manager.ManagerLogin, manager.Password, manager.Name, manager.Surname)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthSQLServer) GetUser(managerLogin, password string) (models.Manager, error) {
	var user models.Manager
	query := fmt.Sprintf("SELECT ManagerID FROM %s WHERE ManagerLogin='%s' AND Password='%s'", managerTable, managerLogin, password)
	row := r.db.QueryRow(query)
	err := row.Scan(&user.ManagerID)

	return user, err
}
