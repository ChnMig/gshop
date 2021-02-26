package user

import (
	"gshop/db"
)

// useNickNameGetUserFromDB Find the user's basic information by nickname
func useNickNameGetUserFromDB(nickName string) (*User, error) {
	u := User{}
	res := db.DB.Where(&User{NickName: nickName}).First(&u)
	if res.Error != nil {
		return nil, res.Error
	}
	return &u, nil
}
