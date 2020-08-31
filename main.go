package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jdxj/words/db"
	"github.com/jdxj/words/logger"
	"github.com/jdxj/words/services"
)

func main() {
	srv := services.NewServer()
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error("Server.ListenAndServe: %s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server.Shutdown: %s", err)
	}
	<-ctx.Done()

	if err := db.Close(); err != nil {
		logger.Error("db.Close: %s", err)
	}
	logger.Info("shutdown")
	_ = logger.Sync()
}
