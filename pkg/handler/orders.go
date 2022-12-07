package handler

import (
	"net/http"
	"strconv"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

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

// !!
func (h *Handler) deleteOrdersManager(c *gin.Context){
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
