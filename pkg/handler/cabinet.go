package handler

import (
	"net/http"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getInfoManager(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	manager, err := h.services.GetInfoManager(managerId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, manager)
}

func (h *Handler) updateManager(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	var manager models.Manager
	if err = c.BindJSON(&manager); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateManager(managerId, manager)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getInfoClient(c *gin.Context){
	clienId, err := getClientId(c)
	if err != nil {
		return
	} 

	client, err := h.services.GetInfoClient(clienId)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *Handler) updateClient(c *gin.Context){
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	var client models.Client
	if err = c.BindJSON(&client); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateClient(clientId, client)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, statusResponse{"ok"})
}