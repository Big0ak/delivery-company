package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

}