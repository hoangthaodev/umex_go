package initialize

import (
	"fmt"

	"go.uber.org/zap"
	"umex.com/global"
)

func Run() {
	// load configuration and initialize application
	LoadConfig()
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("Logger Start", "Success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	serverPort := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverPort)
}
