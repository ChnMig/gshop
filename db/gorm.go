package db

import (
	"gshop/tools"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB database
var DB *gorm.DB

// ConnDB Connect to MySQL database
func ConnDB(dsn string) (err error) {
	// init db session
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		tools.Log.Panic("An error occurred while initializing the db session")
		return err
	}
	return nil
}
