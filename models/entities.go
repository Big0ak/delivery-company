package models

import "database/sql"

type Auto struct {
	AutoId      sql.NullInt64  `json:"id"`
	DriverID    int            `json:"driverId"`
	Model       string         `json:"model" binding:"required"`
	Number      string         `json:"number" binding:"required"`
	Capacity    int            `json:"capacity" binding:"required"`
	Description sql.NullString `json:"description"`
	YearRelease int            `json:"yearRelease" binding:"required"`
}

type Driver struct {
	DriverID int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
}

type Orders struct {
	OrderID     int           `json:"id"`
	ClientID    int           `json:"clientId" binding:"required"`
	DriverID    int           `json:"driverid" binding:"required"`
	ManagerID   sql.NullInt64 `json:"managerId"`
	CargoWeight int           `json:"cargoWeight" binding:"required"`
	Price       uint          `json:"price" binding:"required"`
	Departure   string        `json:"departure" binding:"required"`
	Destination string        `json:"destination" binding:"required"`
	Date        string        `json:"date"`
}

type OrdersRead struct {
	OrderID     int            `json:"id"`
	Client      string         `json:"client" binding:"required"`
	Driver      string         `json:"driver" binding:"required"`
	Manager     sql.NullString `json:"manager" binding:"required"`
	CargoWeight int            `json:"cargoWeight" binding:"required"`
	Price       uint           `json:"price" binding:"required"`
	Departure   string         `json:"departure" binding:"required"`
	Destination string         `json:"destination" binding:"required"`
	Date        string         `json:"date"`
}