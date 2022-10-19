package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autharizationHeader = "Authorization"
	managerCtx = "managerId"
)

// получать значение из header авторизации, валидировать его,
// парсить токен и записывать пользователя в контекст
func (h *Handler) managerIdentity(c *gin.Context){
	header := c.GetHeader(autharizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header") // 401 пользователь не авторизирован
		return
	}
	
	// парсинг токена 
	managerId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// если операция успешна записать значение id в контекст
	// чтобы иметь доступ к id пользователя который делает запрос в последующих обработчиках
	// которые вызываются после данной прослойки
	c.Set(managerCtx, managerId)
}