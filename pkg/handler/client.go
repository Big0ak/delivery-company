package handler

import (
	"net/http"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

type getAllClientResponce struct {
	Data []models.Client `json:"data"`
}

func (h *Handler) getAllClient(c *gin.Context){
	managerid, err := getManagerId(c)
	if err != nil {
		return
	}

	listClient, err := h.services.GetAllClient(managerid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllClientResponce{
		Data: listClient,
	})
}