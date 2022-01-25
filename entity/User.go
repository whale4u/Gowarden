package entity

import orm "Gowarden/common"

type User struct {
	UserName  string `form:"username" binding:"required,min=3,max=10"`
	Password  string `form:"password" binding:"required"`
	Phone     string `form:"phone" binding:"required"`
	Role      string
	AvatarUrl string
	CreateAt  string
	UpdateAt  string
}

var users []User

func (user *User) Create() (result string, err error) {
	res := orm.DB.Create(&user)
	if res.Error != nil {
		err = res.Error
	}
	return
}

func (user *User) Delete() (Result User, err error) {
	if err = orm.DB.Where("user_name = ?", user.UserName).Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
