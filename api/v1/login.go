package v1

import (
	"Gowarden/dto"
	"Gowarden/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		fmt.Println("orm verify failed!")
		fmt.Println(loginDto)
		return
	}
	if loginDto.UserName == "admin" && loginDto.Password == "pass" {
		hmacUser := utils.HmacUser{
			Username: loginDto.UserName,
		}
		if token, err := utils.GenerateToken(hmacUser); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "true",
				"data": token,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "false",
			"data": "Generate token failed.",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "false",
			"data": loginDto,
		})
	}
}
