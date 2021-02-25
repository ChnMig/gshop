package user

import (
	"gshop/db"
)

// Get user information from the database
func useNickNamegetUserFromDB(nickName string) (*User, error) {
	u := User{}
	res := db.DB.Where(&User{NickName: nickName}).First(&u)
	if res.Error != nil {
		return nil, res.Error
	}
	return &u, nil
}
