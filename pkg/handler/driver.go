package handler

import (
	"net/http"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

type getAllDriverResponce struct {
	Data []models.Driver `json:"data"`
}

func (h *Handler) getAllDriver(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	listDriver, err := h.services.GetAllDriver(managerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllDriverResponce{
		Data: listDriver,
	})
}
