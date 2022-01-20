package main

import (
	"Gowarden/config"
	"Gowarden/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	Engine := gin.Default()
	router.InitRouter(Engine)
	Engine.Run(config.PORT)
}
