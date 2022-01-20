package dto

type LoginDto struct {
	UserName string `form:"username" binding:"required,min=3,max=10"`
	Password string `form:"password" binding:"required"`
}
