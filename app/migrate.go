package app

import (
	"gshop/app/user"
	"gshop/db"
)

// MigrateAll Migrate all table structures to mysql database
func MigrateAll() {
	db.DB.AutoMigrate(&user.User{})
}
