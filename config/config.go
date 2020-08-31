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
	key, port := "port", "8080"
	viper.SetDefault(key, port)
	return viper.GetString(key)
}

func GetMode() string {
	key, mode := "mode", gin.DebugMode
	viper.SetDefault(key, mode)
	return viper.GetString(key)
}

func GetSQLite() string {
	key, val := "sqlite", "words.db"
	viper.SetDefault(key, val)
	return viper.GetString(key)
}

func GetLogPath() string {
	key, path := "log_path", "words.log"
	viper.SetDefault(key, path)
	return viper.GetString(key)
}

func GetSecret() []byte {
	key := "secret"
	return []byte(viper.GetString(key))
}

func GetTranslator() string {
	key, val := "translator", "google"
	viper.SetDefault(key, val)
	return viper.GetString(key)
}

func GetDatabase() string {
	key, val := "database", "mysql"
	viper.SetDefault(key, val)
	return viper.GetString(key)
}

type MySQL struct {
	User string
	Pass string
	Addr string
	Base string
}

func GetMySQL() (mysql MySQL) {
	key := "mysql"
	_ = viper.UnmarshalKey(key, &mysql)
	return
}
