package handler

import (
	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
	//"github.com/Big0ak/DeliveryCompany/pkg/service"
)

type Handler struct{
	//services *service.Service
	services Services
}

///////////////////////////////////////////////////////////////////////////////// 
type Authorization interface {
	CreateManagr(models.Manager) (int, error)
	GenerateToken(managerLogin, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Orders interface {
	Create(managerId int, order models.Orders) (int, error)
	GetAll(managerId int) ([]models.Orders, error)
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.managerIdentity)
	{
		orders := api.Group("/orders")
		{
			orders.POST("/", h.createOrdersManager)
			orders.GET("/", h.getAllOrders)
			orders.GET("/:id", h.getOrdersById)
			orders.PUT("/:id", h.updateOrders)
			orders.DELETE("/:id", h.deleteOrders)
		}
		
		route := api.Group("/route")
		{
			route.POST("/", h.createRoute)
			route.GET("/", h.getAllRoute)
			route.GET("/:id", h.getRouteById)
			route.PUT("/:id", h.updateRoute)
			route.DELETE("/:id", h.deleteRoute)
		}
	}
	return router
}