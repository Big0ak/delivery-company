package handler

import (
	"net/http"

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

	id, err := h.services.Create(managerId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"OrderId": id,
	})
}

type getAllOrdersResponse struct {
	Data[] models.Orders `json:"data"`
}

func (h *Handler) getAllOrders(c *gin.Context){
	managerid, err := getManagerId(c)
	if err != nil {
		return
	}

	listOrsers, err := h.services.GetAll(managerid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON (http.StatusOK, getAllOrdersResponse {
		Data: listOrsers,
	})

}

func (h *Handler) getOrdersById(c *gin.Context){

}

func (h *Handler) updateOrders(c *gin.Context){

}

func (h *Handler) deleteOrders(c *gin.Context){

}