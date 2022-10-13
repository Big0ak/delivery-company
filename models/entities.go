package models

type Auto struct {
	AutoId      int    `json:"id"`
	DriverID    int    `json:"driverId"`
	Model       string `json:"model"`
	Number      string `json:"number"`
	Capacity    int    `json:"capacity"`
	Description string `json:"description"`
	YearRelease int    `json:"yearRelease"`
}

type Driver struct {
	DriverID int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type Route struct {
	RouteID     int    `json:"id"`
	Departure   string `json:"departure"`
	Destination string `json:"destination"`
	Distance    int    `json:"distance"`
}

type Orders struct {
	OrderID     int    `json:"id"`
	ClientID    int    `json:"clientId"`
	RouteID     int    `json:"routeId"`
	DriverID    int    `json:"driverid"`
	ManagerID   int    `json:"managerId"`
	CargoWeight int    `json:"cargoWeight"`
	Price       int    `json:"price"`
	Date        string `json:"date"`
}
