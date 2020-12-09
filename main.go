package gshop

import (
	"io"
	"os"

	"github.com/chnmig/gshop/app"
	"github.com/chnmig/gshop/db"
	"github.com/chnmig/gshop/tools"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
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
		// gin log to file
		ginFile, _ := os.OpenFile("./gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		gin.DefaultWriter = io.MultiWriter(ginFile)
	}
	// run gin
	r := app.InitApp()
	err := r.Run(tools.EnvConfig.Gshop.HOST)
	if err != nil {
		tools.Log.Panic("start gin error", zap.Error(err))
	}
	tools.Log.Info("successfully started Gin")
}
