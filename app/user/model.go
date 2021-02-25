package user

import "gorm.io/gorm"

type nickNameLoginParam struct {
	NickName string `json:"nick_name" from:"nick_name" binding:"required" max:"10" min:"4"`
	PassWord string `json:"pass_word" from:"pass_word" binding:"required" max:"20" min:"6"`
}

type nameLoginResponse struct {
	Token string `json:"token"`
}

// User user model from grom
type User struct {
	gorm.Model
	NickName string `gorm:"size:10"`
	PassWord string `gorm:"size:20"`
}
