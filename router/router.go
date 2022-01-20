package router

import (
	"Gowarden/api/v1"
	"Gowarden/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.Engine) {
	Router.POST("/login", v1.Login)
	GroupV1 := Router.Group("/v1").Use(middleware.AuthMiddleWare())
	{
		GroupV1.POST("/home", v1.Home)
	}
}
