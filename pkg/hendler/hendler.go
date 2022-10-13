package handler

import (
	"github.com/gin-gonic/gin"
	//"github.com/Big0ak/DeliveryCompany/pkg/service"

)

type Handler struct{
	//services *service.Service
	services Services

}

///////////////////////////////////////////////////////////////////////////////// 

type Services interface {

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

	api := router.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.POST("/", h.createOrders)
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