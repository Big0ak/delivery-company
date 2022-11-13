package handler

import (
	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
	//"github.com/Big0ak/DeliveryCompany/pkg/service"
	"github.com/gin-contrib/cors"
)

type Handler struct{
	//services *service.Service
	services Services
}

///////////////////////////////////////////////////////////////////////////////// 
type Authorization interface {
	CreateNewManager(models.Manager) (int, error)
	GenerateTokenManager(managerLogin, password string) (string, error)
	ParseTokenManager(token string) (int, error)
	
	CreateNewClient(models.Client) (int, error)
	GenerateTokenClient(managerLogin, password string) (string, error)
	ParseTokenClient(token string) (int, error)
}

type Orders interface {
	CreateManager(managerId int, order models.Orders) (int, error)
	GetAll(managerId int) ([]models.Orders, error)
	GetByIdManager(managerid, id int) (models.Orders, error)
	DeleteManager(managerid, id int) error
	UpdateManager(managerid, id int, input models.Orders) error
}

type Services interface {
	Authorization
	Orders
}

func NewHandler (services Services) *Handler{
	return &Handler{services: services}
}

// func NewHandler(services *service.Service) *Handler {
// 	return &Handler{services: services}
// }
///////////////////////////////////////////////////////////////////////////////////
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
		// регистрация и авторизация менеджера
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

		// регистрация и авторизация пользователя
		auth.POST("/client-sign-up", h.clientSignUp)
		auth.POST("/client-sign-in", h.clientSignIn)
	}

	manager := router.Group("/manager-api", h.managerIdentity)
	{
		orders := manager.Group("/orders")
		{
			orders.POST("/", h.createOrdersManager)
			orders.GET("/", h.getAllOrders)
			orders.GET("/:id", h.getOrdersById)
			orders.PUT("/:id", h.updateOrders)
			orders.DELETE("/:id", h.deleteOrdersManager)
		}
	}

	user := router.Group("/client-api", h.clientIdentity)
	{
		orders := user.Group("/orders")
		{
			orders.GET("/", h.getUserOrder)
		}
	}

	return router
}