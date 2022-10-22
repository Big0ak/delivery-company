package models

// binding:"required" валидация для Gin
type Client struct {
	ClientID         int    `json:"-"`
	UserLogin        string `json:"login" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Surname          string `json:"surname" binding:"required"`
	Phone            string `json:"phone" binding:"required"`
	RegistrationDate string `json:"registrationDate"`
}

type Manager struct {
	ManagerID    int    `json:"-"`
	ManagerLogin string `json:"login" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Surname      string `json:"surname" binding:"required"`
}
