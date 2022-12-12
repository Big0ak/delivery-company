package handler

import (
	"net/http"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

type getAllDriverResponce struct {
	Data []models.Driver `json:"data"`
}

// @Summary Get all driver
// @Security ApiKeyAuth
// @Tags Manager
// @Description Get all driver by a manager
// @ID get-all-driver-manager
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/driver/ [get]
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
