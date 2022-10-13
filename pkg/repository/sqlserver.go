package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

type Config struct {
	Server   string
	User     string
	Password string
	Port     string
	Database string
}

func GetDB(cfg Config) (db *sql.DB, err error){
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		cfg.Server, cfg.User, cfg.Password, cfg.Port, cfg.Database)
	
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil{
		return nil, err
	}

	return db, err
}