package handler

import (
	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	services Services
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type Authorization interface {
	CreateNewManager(models.Manager) (int, error)
	GenerateTokenManager(managerLogin, password string) (string, error)
	ParseTokenManager(token string) (int, error)
	
	CreateNewClient(models.Client) (int, error)
	GenerateTokenClient(managerLogin, password string) (string, error)
	ParseTokenClient(token string) (int, error)
}

type Orders interface {
	// функции для менеджера
	CreateManager(managerId int, order models.Orders) (int, error)
	GetAll(managerId int) ([]models.OrdersRead, error)
	GetAllWithStatus(managerId int, status string) ([]models.OrdersRead, error)
	DeleteManager(managerId, id int) error
	UpdateOrderManager(managerId, id int, input models.Orders) error
	SearchOrdersByCityManager(managerId int, city string) ([]models.OrdersRead, error)
	
	// для менеджера и клиента
	GetById(userId, id int) (models.OrdersRead, error)

	// для клиента
	GetUserOrder(clientId int) ([]models.OrdersRead, error)
	GetAllWithStatusUserDB(clientId int, status string) ([]models.OrdersRead, error)
	SearchOrdersByCityClient(clientId int, city string) ([]models.OrdersRead, error) 
}

type Client interface {
	GetAllClient(managerId int) ([]models.Client, error)
}

type Driver interface {
	GetAllDriver(managerId int) ([]models.Driver, error)
}

type Cabinet interface {
	// функции менеджера
	GetInfoManager(managerId int) (models.Manager, error)
	UpdateManager(mamagerId int, manager models.Manager) error 

	// функции клиента
	GetInfoClient(clientId int) (models.Client, error)
	UpdateClient(clientId int, client models.Client) error
}

type Services interface {
	Authorization
	Orders
	Client
	Driver
	Cabinet
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func NewHandler (services Services) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization"},
		AllowCredentials: true,
	}))

	auth := router.Group("/auth")
	{
		// регистрация  менеджера
		auth.POST("/sign-up", h.signUp)
		
		// регистрация пользователя
		auth.POST("/client-sign-up", h.clientSignUp)

		// авторизация клиента и менеджера
		auth.POST("/sign-in", h.signIn)
	}

	manager := router.Group("/manager-api", h.managerIdentity)
	{
		orders := manager.Group("/orders")
		{
			orders.POST("/", h.createOrdersManager)
			orders.GET("/", h.getAllOrders)
			orders.GET("/active", h.getAllActiveOrders)
			orders.GET("/completed", h.getAllCompletedOrders)
			orders.GET("/:id", h.getOrderByIdManager)
			orders.PUT("/:id", h.updateOrders)
			orders.DELETE("/:id", h.deleteOrdersManager)
			orders.GET("/search/:city", h.searchOrdersByCityManager)
		}

		client := manager.Group("/client")
		{
			client.GET("/", h.getAllClient)
		}

		driver := manager.Group("/driver")
		{
			driver.GET("/", h.getAllDriver)
		}

		cabinet := manager.Group("/cabinet")
		{
			cabinet.GET("/", h.getInfoManager)
			cabinet.PUT("/", h.updateManager)
		}
	}

	user := router.Group("/client-api", h.clientIdentity)
	{
		orders := user.Group("/orders")
		{
			orders.GET("/", h.getUserOrder)
			orders.GET("/active", h.getAllActiveUserOrders)
			orders.GET("/completed", h.getAllCompletedUserOrders)
			orders.GET("/:id", h.getOrderByIdClient) //Возможно лучше делать отдельным запросом
			orders.GET("/search/:city", h.searchOrdersByCityClient)
		}

		cabinet := user.Group("/cabinet")
		{
			cabinet.GET("/", h.getInfoClient)
			cabinet.PUT("/", h.updateClient)
		}
	}

	return router
}