package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(ctx *gin.Context) {
	Name, _ := ctx.Get("username")
	UserName := Name.(string)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "true",
		"data": UserName,
	})
}
