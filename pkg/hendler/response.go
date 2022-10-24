package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statsCode int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(statsCode, errorResponse{message})
}