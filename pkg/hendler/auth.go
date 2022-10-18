package handler

import (
	"net/http"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.Manager

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 некорректные данные
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 ошибка на сервере
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	ManagerLogin string `json:"login" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 некорректные данные
		return
	}

	token, err := h.services.GenerateToken(input.ManagerLogin, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 ошибка на сервере
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
