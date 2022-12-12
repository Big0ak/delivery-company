package handler

import (
	"net/http"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

type getAllClientResponce struct {
	Data []models.Client `json:"data"`
}

// @Summary Get all client
// @Security ApiKeyAuth
// @Tags Manager
// @Description Get all clietn by a manager
// @ID get-all-client-manager
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/client/ [get]
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