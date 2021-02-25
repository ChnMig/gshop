package main

import (
	"flag"
	"fmt"
	"gshop/app"
	"gshop/db"
	tools "gshop/tool"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title Gshop API
// @description Gshop is an API for shopping websites
// @BasePath /api
// @contact.name ChnMig
// @contact.email ChnMig@Outlook.com
// @version 0.1
// @license.name Apache 2.0
// @license.url https://github.com/ChnMig/gshop/blob/main/LICENSE

func main() {
	op := flag.String("operation", "start", `This determines the operation of the program
	start: Start project, is default
	migrate: Synchronize database structure to database
	init: Add basic data to database
	`)
	switch *op {
	case "start":
		start()
	case "migrate":
		migrate()
	default:
		fmt.Println("Unexpected operation")
	}
}

// start project
func start() {
	flag.Parsed()
	// init gshop log
	tools.InitLogger()
	// Clear the buffer before the end of the program
	defer tools.Log.Sync()
	tools.Log.Info("successfully start logger")
	// get env
	tools.DoEnv()
	// init config
	tools.InitConfig()
	tools.Log.Info("successfully get env")
	// init grom
	_ = db.ConnDB(tools.EnvConfig.DB.Address)
	tools.Log.Info("successfully init gorm")
	// init redis
	_ = db.ConnRedis(tools.EnvConfig.Redis.Address)
	defer db.RDB.Close()
	tools.Log.Info("successfully init redis")
	// start gshop
	if tools.EnvConfig.Gshop.Debug == 0 {
		// close gin debug
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
	// run gin
	r := app.InitApp()
	err := r.Run(tools.EnvConfig.Gshop.HOST)
	if err != nil {
		tools.Log.Panic("start gin error", zap.Error(err))
	}
	tools.Log.Info("successfully started Gin")
}

func migrate() {
	// init config
	tools.InitConfig()
	tools.Log.Info("successfully get env")
	// init grom
	_ = db.ConnDB(tools.EnvConfig.DB.Address)
	tools.Log.Info("successfully init gorm")
	// migrate
	app.MigrateAll()
	fmt.Println("successfully")
}
