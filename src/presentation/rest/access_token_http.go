package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/johnwoz123/pharmacy-authorization-api/src/domain/access_token"
	"net/http"
)

// This is the blue layer of the onion - controller

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHttpHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := c.Param("access_token_id")
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, accessToken)
}
