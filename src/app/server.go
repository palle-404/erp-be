package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/palle-404/erp-be/src/config"
	"github.com/palle-404/erp-be/src/db"
	"github.com/palle-404/erp-be/src/logger"
	"github.com/palle-404/erp-be/src/service"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var quit = make(chan os.Signal, 1)
var errServerStart = make(chan bool)

var server *http.Server

func Start() error {
	if err := logger.Init(); err != nil {
		return err
	}
	logger.Log().Info("Logger initialized ...")
	ctx := context.TODO()
	if err := db.Connect(ctx); err != nil {
		logger.Log().Error("Failed to connect to database", zap.Error(err))
		return err
	}
	dbLayer := db.NewDBLayer()
	apiLayer := service.NewServiceLayer(dbLayer)

	port := config.AppCfg().GetInt("app.port")
	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: createRouter(apiLayer),
	}
	logger.Log().Info("Starting server", zap.Int("port", port))
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log().Error("Server start error", zap.Error(err))
			errServerStart <- true
		}
	}()
	return nil
}

func ListenForShutdown() {
	select {
	case <-quit:
	case <-errServerStart:
	}
	logger.Log().Info("Server shutdown initiated ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Log().Error("Server shutdown error", zap.Error(err))
	}
	if err := db.Disconnect(ctx); err != nil {
		logger.Log().Error("Database disconnect error", zap.Error(err))
	}
	if err := logger.Log().Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
		log.Println("Logger sync error" + err.Error())
	}
	log.Fatal("Server shutdown gracefully.")
}
