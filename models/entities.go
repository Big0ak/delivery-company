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

type Route struct {
	RouteID     int    `json:"id"`
	Departure   string `json:"departure" binding:"required"`
	Destination string `json:"destination" binding:"required"`
	Distance    int    `json:"distance" binding:"required"`
}

type Orders struct {
	OrderID     int           `json:"id"`
	ClientID    int           `json:"clientId" binding:"required"`
	RouteID     int           `json:"routeId" binding:"required"`
	DriverID    int           `json:"driverid" binding:"required"`
	ManagerID   sql.NullInt64 `json:"managerId"`
	CargoWeight int           `json:"cargoWeight" binding:"required"`
	Price       uint          `json:"price" binding:"required"`
	Date        string        `json:"date"`
}

// type UpdateOrders struct {
// 	ClientId    *int    `json:"id"`
// 	RouteID     *int    `json:"routeId"`
// 	DriverID    *int    `json:"driverid"`
// 	ManagerID   *int    `json:"managerId"`
// 	CargoWeight *int    `json:"cargoWeight"`
// 	Price       *uint   `json:"price"`
// 	Date        *string `json:"date"`
// }
