package main

import (
	"fmt"

	"github.com/jdxj/words/config"
	"github.com/jdxj/words/logger"
	"github.com/jdxj/words/services"
)

func main() {
	// todo: 关闭 DB

	bind := fmt.Sprintf(":%s", config.GetPort())
	router := services.NewRouter()
	if err := router.Run(bind); err != nil {
		logger.Error("router.Run: %s", err)
	}
}
