package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autharizationHeader = "Authorization"
	managerCtx          = "managerId"
	clientCtx			= "clientId"
)

// получать значение из header авторизации, валидировать его,
// парсить токен и записывать пользователя в контекст
func (h *Handler) managerIdentity(c *gin.Context) {
	header := c.GetHeader(autharizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header") // 401 менеджер не авторизирован
		return
	}

	// парсинг токена
	managerId, err := h.services.ParseTokenManager(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// если операция успешна записать значение id в контекст
	// чтобы иметь доступ к id пользователя который делает запрос в последующих обработчиках
	// которые вызываются после данной прослойки
	c.Set(managerCtx, managerId)
}

func (h *Handler)clientIdentity(c *gin.Context){
	header:= c.GetHeader(autharizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header") // 401 клиент не авторизирован
		return
	}

	// парсинг токена
	clientId, err := h.services.ParseTokenClient(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(clientCtx, clientId)
}

// получение Id менеджера, чтобы каждый раз не прописывать
func getManagerId(c *gin.Context) (int, error) {
	id, ok := c.Get(managerCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "manager Id not found")
		return 0, errors.New("manager Id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "manager Id is of invalid type")
		return 0, errors.New("manager Id is of invalid type")
	}
	return idInt, nil
}

func getClientId(c *gin.Context) (int, error) {
	id, ok := c.Get(clientCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "client Id not found")
		return 0, errors.New("client Id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "client Id is of invalid type")
		return 0, errors.New("client Id is of invalid type")
	}
	return idInt, nil
}