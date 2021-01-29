package app

// 总路由

import (
	"github.com/chnmig/gshop/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitApp init gshop app
func InitApp() *gin.Engine {
	// gin.Default uses Use by default. Two global middlewares are added, Logger(), Recovery(), Logger is to print logs, Recovery is panic and returns 500
	router := gin.Default()
	// gin log
	// router.Use(tools.GinLogger(), tools.GinRecovery(true))
	// Add consent cross-domain middleware
	router.Use(middleware.CorsHandler())

	// v1
	router.Group("/api/v1")
	{

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
