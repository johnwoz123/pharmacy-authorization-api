package app

import (
	"github.com/gin-gonic/gin"
	"github.com/johnwoz123/pharmacy-authorization-api/src/domain/access_token"
	"github.com/johnwoz123/pharmacy-authorization-api/src/presentation/rest"
	"github.com/johnwoz123/pharmacy-authorization-api/src/repository/db/mysql"
)

var (
	router = gin.Default()
)

// handler - Presentation Layer - Blue Layer
// service - Interface Adapter - Green Layer
// repository -
func StartApp() {
	service := access_token.NewService(mysql.NewRepo())
	handler := rest.NewHttpHandler(service)

	router.POST("/auth/access_token", handler.Create)
	router.GET("/auth/access_token/:access_token_id", handler.GetById)
	router.Run(":8081")
}
