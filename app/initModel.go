package app

import (
	"gshop/db"
	"os/user"
)

// MigrateAll Migrate all table structures to mysql database
func MigrateAll() {
	db.DB.AutoMigrate(&user.User{})
}
