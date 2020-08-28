package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "yaml"
	// todo: 路径更改
	configPath = "/home/jdxj/workspace/words/config"
)

const (
	portKey = "port"
	modeKey = "mode"
)

func init() {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetPort() string {
	port := "8080"
	viper.SetDefault(portKey, port)
	return viper.GetString(portKey)
}

func GetMode() string {
	viper.SetDefault(modeKey, gin.DebugMode)
	return viper.GetString(modeKey)
}
