package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"umex.com/global"
)

func LoadConfig() {
	// read file by viper
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read config file: %w", err))
	}

	// read server configuration
	fmt.Println("server port: ", viper.GetInt("server.port"))

	// configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config struct: %v", err)
	}
}
