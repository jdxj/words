package main

import (
	"github.com/jdxj/words/services"
	"go.uber.org/zap"
)

func main() {
	// todo: 关闭 DB

	router := services.NewRouter()
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}

func initLogger() {
	zap.NewDevelopment()
}
