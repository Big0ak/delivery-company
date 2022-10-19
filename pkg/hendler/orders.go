package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createOrders(c *gin.Context){
	id, _ := c.Get(managerCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id" : id,
	})
}

func (h *Handler) getAllOrders(c *gin.Context){

}

func (h *Handler) getOrdersById(c *gin.Context){

}

func (h *Handler) updateOrders(c *gin.Context){

}

func (h *Handler) deleteOrders(c *gin.Context){

}