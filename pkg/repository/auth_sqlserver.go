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

func (r *AuthSQLServer) CreateNewManagerDB(manager models.Manager) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (ManagerLogin, Password, Name, Surname) OUTPUT Inserted.ManagerID VALUES ('%s', '%s', '%s', '%s')", managerTable, manager.ManagerLogin, manager.Password, manager.Name, manager.Surname)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthSQLServer) CreateNewClientDB(client models.Client) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (UserLogin, Password, Name, Surname, Phone) OUTPUT Inserted.ClientID VALUES ('%s', '%s', '%s', '%s', '%s')",
		clientTable, client.UserLogin, client.Password, client.Name, client.Surname, client.Phone)
	row := r.db.QueryRow(query)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthSQLServer) GetManager(managerLogin, password string) (models.Manager, error) {
	var manager models.Manager
	query := fmt.Sprintf("SELECT ManagerID FROM %s WHERE ManagerLogin='%s' AND Password='%s'", managerTable, managerLogin, password)
	row := r.db.QueryRow(query)
	err := row.Scan(&manager.ManagerID)

	return manager, err
}

func (r *AuthSQLServer) GetClient(clientLogin, password string) (models.Client, error) {
	var client models.Client
	query := fmt.Sprintf("SELECT ClientID FROM %s WHERE UserLogin='%s' AND Password='%s'", clientTable, clientLogin, password)
	row := r.db.QueryRow(query)
	err := row.Scan(&client.ClientID)

	return client, err
}
