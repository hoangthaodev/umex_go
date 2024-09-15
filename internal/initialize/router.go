package initialize

import (
	"github.com/gin-gonic/gin"
	"umex.com/global"
	"umex.com/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middleware
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limitter

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("v1/2024")
	{
		MainGroup.GET("/checkStatus") // tracking status
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}

	return r
}
