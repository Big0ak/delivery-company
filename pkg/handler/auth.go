package handler

import (
	"net/http"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

// имя, где передается в http (тип параметра), его структура, обязателен(или нет), описание 
// Param input body models.Manager true "account info"

// @Summary Manager SignUp
// @Tags Auth
// @Description create manager account
// @ID create-manager-account
// @Accept  json
// @Produce  json 
// @Param input body models.Manager true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.Manager

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 некорректные данные
		return
	}

	id, err := h.services.CreateNewManager(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 ошибка на сервере
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}


// @Summary Client SignUp
// @Tags Auth
// @Description creat client account
// @ID creat-client-account
// @Accept json
// @Produce json
// @Param input body models.Client true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/client-sign-up [post]
func (h *Handler) clientSignUp(c *gin.Context) {
	var input models.Client

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400
		return
	}

	id, err := h.services.CreateNewClient(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type signInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags Auth
// @Description login to your personal account
// @ID login-account
// @Accept json
// @Produce json
// @Param input body signInInput true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // 400 некорректные данные
		return
	}

	token, err := h.services.GenerateTokenClient(input.Login, input.Password)
	if err != nil {
		token, err = h.services.GenerateTokenManager(input.Login, input.Password)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error()) // 500 ошибка на сервере
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
				"role": "manager",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"role": "client",
		})
	}	
}
