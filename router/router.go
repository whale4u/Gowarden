package router

import (
	"Gowarden/middleware"
	v1 "Gowarden/router/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.Engine) {
	Router.POST("/login", v1.Login)
	Router.POST("/register", v1.AddUser)
	Router.POST("/delete", v1.DelUser)
	GroupV1 := Router.Group("/v1").Use(middleware.AuthMiddleWare())
	{
		GroupV1.POST("/home", v1.Home)
	}
}
