package models

type Client struct {
	ClientID         int    `json:"-"`
	UserLogin        string `json:"login"`
	Password         string `json:"password"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	Phone            string `json:"phone"`
	RegistrationDate string `json:"registrationDate"`
}

type Manager struct {
	ManagerID    int    `json:"-"`
	ManagerLogin string `json:"login"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
}
