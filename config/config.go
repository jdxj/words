package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "yaml"

	// 硬编码
	configPathLocal = "."
	configPathDebug = "/home/jdxj/workspace/words"
)

const (
	portKey    = "port"
	modeKey    = "mode"
	dbPathKey  = "db_path"
	logPathKey = "log_path"
	secretKey  = "secret"

	translatorKey = "translator"
)

func init() {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPathLocal)
	viper.AddConfigPath(configPathDebug)

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

func GetDBPath() string {
	viper.SetDefault(dbPathKey, "words.db")
	return viper.GetString(dbPathKey)
}

func GetLogPath() string {
	viper.SetDefault(logPathKey, "words.log")
	return viper.GetString(logPathKey)
}

func GetSecret() []byte {
	viper.SetDefault(secretKey, "")
	return []byte(viper.GetString(secretKey))
}

func GetTranslator() string {
	viper.SetDefault(translatorKey, "google")
	return viper.GetString(translatorKey)
}
