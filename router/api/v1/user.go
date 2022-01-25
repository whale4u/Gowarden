package v1

import (
	"Gowarden/entity"
	"Gowarden/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(ctx *gin.Context) {
	var user entity.User

	//从HTTP请求包中解析出json
	var params map[string]string
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.Decode(&params)

	user.UserName = params["username"]
	plainText := params["password"]
	user.Password, _ = utils.GeneratePassword(plainText)
	user.Phone = params["phone"]
	user.Role = "user"
	user.AvatarUrl = "www.sina.com"
	user.CreateAt = "2022-01-23 21:53:39"
	user.UpdateAt = "2022-01-23 21:54:27"

	// 用户名不可为空
	if user.UserName == "" {
		fmt.Println("username: ", user.UserName)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "true",
			"data": "UserName can't be empty.",
		})
		return
	}
	// 注册用户
	_, err := user.Create()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "true",
			"data": "User register error.",
		})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "true",
		"data": "User register ok.",
	})
}

func DelUser(ctx *gin.Context) {
	var user entity.User
	//从HTTP请求包中解析出json
	var params map[string]string
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.Decode(&params)

	user.UserName = params["username"]
	// 用户名不可为空
	if user.UserName == "" {
		fmt.Println("username: ", user.UserName)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "true",
			"data": "UserName can't be empty.",
		})
		return
	}

	//缺少用户查询逻辑

	_, err := user.Delete()
	if err != nil {
		return
	}
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "true",
			"data": "User delete error.",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "true",
		"data": "User delete ok.",
	})
}
