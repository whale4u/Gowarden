package middleware

import (
	"Gowarden/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("token")
		fmt.Println("current token: ", tokenString)
		if tokenString == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "false",
				"data": "token is null.",
			})
			context.Abort()
			return
		}
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "false",
				"data": "token parse failed.",
			})
			context.Abort()
			return
		}
		context.Set("username", claims.Username)
		context.Next()
	}
}
