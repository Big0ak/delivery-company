package handler

import (
	"net/http"
	"strconv"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

// @Summary Creat order
// @Security ApiKeyAuth
// @Tags Manager
// @Description creating an order by a manager
// @ID creat-order-manager
// @Accept json
// @Produce json
// @Param input body models.Orders true "order info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/orders/ [post]
func (h *Handler) createOrdersManager(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	var input models.Orders
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateManager(managerId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"OrderId": id,
	})
}

type getOrdersResponse struct {
	Data[] models.OrdersRead `json:"data"`
}

// @Summary Get all orders
// @Security ApiKeyAuth
// @Tags Manager
// @Description Get all orders by a manager
// @ID get-all-orders-manager
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/orders/ [get]
func (h *Handler) getAllOrders(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	listOrsers, err := h.services.GetAll(managerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrdersResponse {
		Data: listOrsers,
	})
}

// @Summary Get all active orders
// @Security ApiKeyAuth
// @Tags Manager
// @Description Get all active orders by a manager
// @ID get-all-active-orders-manager
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/orders/active [get]
func (h *Handler) getAllActiveOrders(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	listOrders, err := h.services.GetAllWithStatus(managerId, "active")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrdersResponse{
		Data: listOrders,
	})
}

// @Summary Get all completed orders
// @Security ApiKeyAuth
// @Tags Manager
// @Description Get all completed orders by a manager
// @ID get-all-completed-orders-manager
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/orders/completed [get]
func (h *Handler) getAllCompletedOrders(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	listOrders, err := h.services.GetAllWithStatus(managerId, "completed")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrdersResponse{
		Data: listOrders,
	})
}

// @Summary Get order by id
// @Security ApiKeyAuth
// @Tags Manager
// @Description get order by id manager
// @ID get-order-byid-manager
// @Accept json
// @Produce json
// @Param id path int true "ordre id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/orders/{id} [get]
func (h *Handler) getOrderByIdManager(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	order, err := h.services.GetById(managerId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON (http.StatusOK, order)
}

// @Summary Update order
// @Security ApiKeyAuth
// @Tags Manager
// @Description update order by a manager
// @ID update-order-manager
// @Accept json
// @Produce json
// @Param id path int true "order id"
// @Param input body models.Orders true "new info order"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/orders/{id} [put]
func (h *Handler) updateOrders(c *gin.Context){
	managerid, err := getManagerId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.Orders
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateOrderManager(managerid, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON (http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete order
// @Security ApiKeyAuth
// @Tags Manager
// @Description delete order by a manager
// @ID delete-order-manager
// @Accept json
// @Produce json
// @Param id path int true "order id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/orders/{id} [delete]
func (h *Handler) deleteOrdersManager(c *gin.Context){ // !!
	managerid, err := getManagerId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.DeleteManager(managerid, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// всегда возвращает "ок", исправить
	c.JSON (http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Search order
// @Security ApiKeyAuth
// @Tags Manager
// @Description search order by a manager
// @ID search-order-manager
// @Accept json
// @Produce json
// @Param city path string true "name city"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/orders/search/{city} [get]
func (h *Handler) searchOrdersByCityManager(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	city := c.Param("city")

	listOrders, err := h.services.SearchOrdersByCityManager(managerId, city)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrdersResponse{
		Data: listOrders,
	})
}

// -------------------------------------------------------------------------------------------------
// ------------------------------ Client function --------------------------------------------------
// -------------------------------------------------------------------------------------------------

// @Summary Get all client orders
// @Security ApiKeyAuth
// @Tags Client
// @Description Get all orders by a client
// @ID get-all-orders-client
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /client-api/orders/ [get]
func (h *Handler) getUserOrder (c *gin.Context){
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	listOrsers, err := h.services.GetUserOrder(clientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, getOrdersResponse{
		Data: listOrsers,
	})
}

// @Summary Get active orders
// @Security ApiKeyAuth
// @Tags Client
// @Description Get active orders by a client
// @ID get-all-active-orders-client
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /client-api/orders/active [get]
func (h *Handler) getAllActiveUserOrders (c *gin.Context){
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	listOrders, err := h.services.GetAllWithStatusUserDB(clientId, "active")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrdersResponse{
		Data: listOrders,
	})
}

// @Summary Get completed orders
// @Security ApiKeyAuth
// @Tags Client
// @Description Get completed orders by a client
// @ID get-all-completed-orders-client
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /client-api/orders/completed [get]
func (h *Handler) getAllCompletedUserOrders (c *gin.Context){
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	listOrders, err := h.services.GetAllWithStatusUserDB(clientId, "completed")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrdersResponse{
		Data: listOrders,
	})
}

// @Summary Get order by id
// @Security ApiKeyAuth
// @Tags Client
// @Description get order by id client
// @ID get-order-byid-client
// @Accept json
// @Produce json
// @Param id path int true "ordre id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /client-api/orders/{id} [get]
func (h *Handler) getOrderByIdClient(c *gin.Context){
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	order, err := h.services.GetById(clientId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON (http.StatusOK, order)
}

// @Summary Search order
// @Security ApiKeyAuth
// @Tags Client
// @Description search order by a client
// @ID search-order-client
// @Accept json
// @Produce json
// @Param city path string true "name city"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /client-api/orders/search/{city} [get]
func (h *Handler) searchOrdersByCityClient(c *gin.Context){
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	city := c.Param("city")

	listOrders, err := h.services.SearchOrdersByCityClient(clientId, city)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getOrdersResponse{
		Data: listOrders,
	})
}
