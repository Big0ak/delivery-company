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
	managerid, err := getManagerId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	order, err := h.services.GetByIdManager(managerid, id)
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

	err = h.services.UpdateManager(managerid, id, input)
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

func (h *Handler) getUserOrder (c *gin.Context){

}