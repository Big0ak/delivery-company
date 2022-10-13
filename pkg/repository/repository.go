package repository

import "database/sql"

type Authorization interface{

}

type Orders interface{

}

type Route interface{

}

type Repository struct{
	Authorization
	Orders
	Route
}

func NewRepository (db *sql.DB) *Repository{
	return &Repository{}
}