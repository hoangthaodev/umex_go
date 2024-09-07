package initialize

import (
	"umex.com/global"
	"umex.com/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
