package service

import "github.com/Big0ak/delivery-company/models"

type Service struct {
	*AuthService
	*OrderService
	*ClientService
	*DriverService
	*CabinetService
	repos Repository
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type AuthorizationDB interface{
	CreateNewManagerDB(models.Manager) (int, error)
	GetManager(managerLogin, password string) (models.Manager, error)
	
	CreateNewClientDB(models.Client) (int, error)
	GetClient(clientLogin, password string) (models.Client, error)
}

type OrderDB interface{
	CreateManagerDB(managerId int, order models.Orders) (int, error)
	GetAllDB(managerId int) ([]models.OrdersRead, error)
	DeleteManagerDB(managerId, id int) error
	UpdateManagerDB(managerId, id int, input models.Orders) error
	SearchOrdersByCityManagerDB(managerId int, city string) ([]models.OrdersRead, error)
	
	GetByIdDB(userId, id int) (models.OrdersRead, error)
	GetUserOrderDB(clientId int) ([]models.OrdersRead, error)
	SearchOrdersByCityClientDB(clientId int, city string) ([]models.OrdersRead, error)
}

type ClientDB interface{
	GetAllClientDB(managerId int) ([]models.Client, error)
}

type DriverDB interface{
	GetAllDriverDB(managerId int) ([]models.Driver, error)
}

type CabinetDB interface{
	GetInfoClientDB(clientId int) (models.Client, error)
}

type Repository interface{
	AuthorizationDB
	OrderDB
	ClientDB
	DriverDB
	CabinetDB
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func NewService(repos Repository, AuthDB AuthorizationDB, OrderDB OrderDB, ClientDB ClientDB, DriverDB DriverDB, CabinetDB CabinetDB) *Service {
	return &Service{
		repos: repos,
		AuthService: NewAuthServise(AuthDB),
		OrderService: NewOrderService(OrderDB),
		ClientService: NewClientService(ClientDB),
		DriverService: NewDriverService(DriverDB),
		CabinetService: NewCabinetService(CabinetDB),
	} 
}