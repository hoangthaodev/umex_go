package initialize

import (
	"fmt"

	"umex.com/global"
)

func Run() {
	// load configuration and initialize application
	LoadConfig()
	fmt.Println("Loading configuration...", global.Config.Postgres.Password)
	// initialize logger
	InitLogger()
	global.Logger.Info("Config Log ok!!")
	InitPostgres()

}
