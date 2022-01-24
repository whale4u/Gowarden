package entity

// LoginDto Data Transfer Object（数据传输对象）DTO 是一组需要跨进程或网络边界传输的聚合数据的简单容器。
type LoginDto struct {
	UserName string `form:"username" binding:"required,min=3,max=10"`
	Password string `form:"password" binding:"required"`
}
