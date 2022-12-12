package handler

import (
	"net/http"

	"github.com/Big0ak/delivery-company/models"
	"github.com/gin-gonic/gin"
)

// @Summary Get info manager
// @Security ApiKeyAuth
// @Tags Manager
// @Description getting information about the manager
// @ID get-info-manager
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/cabinet/ [get]
func (h *Handler) getInfoManager(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	manager, err := h.services.GetInfoManager(managerId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, manager)
}

// @Summary update info manager
// @Security ApiKeyAuth
// @Tags Manager
// @Description update information about the manager
// @ID update-info-manager
// @Accept json
// @Produce json
// @Param manager body models.Manager true "new info manager"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /manager-api/cabinet/ [put]
func (h *Handler) updateManager(c *gin.Context){
	managerId, err := getManagerId(c)
	if err != nil {
		return
	}

	var manager models.Manager
	if err = c.BindJSON(&manager); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateManager(managerId, manager)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// -------------------------------------------------------------------------------------------------
// ------------------------------ Client function --------------------------------------------------
// -------------------------------------------------------------------------------------------------

// @Summary Get info client
// @Security ApiKeyAuth
// @Tags Client
// @Description getting information about the client
// @ID get-info-client
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /client-api/cabinet/ [get]
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

// @Summary update info client
// @Security ApiKeyAuth
// @Tags Client
// @Description update information about the client
// @ID update-info-client
// @Accept json
// @Produce json
// @Param client body models.Client true "new info client"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /client-api/cabinet/ [put]
func (h *Handler) updateClient(c *gin.Context){
	clientId, err := getClientId(c)
	if err != nil {
		return
	}

	var client models.Client
	if err = c.BindJSON(&client); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateClient(clientId, client)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, statusResponse{"ok"})
}