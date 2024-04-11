package db

import (
	"context"
	"fmt"
	"github.com/palle-404/erp-be/src/config"
	"github.com/palle-404/erp-be/src/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

var client *mongo.Client

func Connect(ctx context.Context) error {
	cfg := config.AppCfg()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.GetString("db.user"), cfg.GetString("db.password"), cfg.GetString("db.host"), cfg.GetString("db.port"))
	logger.Log().Info("MongoDB URI", zap.String("uri", uri))
	opts := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(ctx, opts)
	if err != nil {
		logger.Log().Error("Could not connect to database", zap.Error(err))
		return err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Log().Error("Could not ping database", zap.Error(err))
		return err
	}
	logger.Log().Info("Connected to database ...")
	loadCollections(client)
	logger.Log().Info("Collections successfully initialized ...")
	return nil
}

func Disconnect(ctx context.Context) error {
	if err := client.Disconnect(ctx); err != nil {
		logger.Log().Error("Could not disconnect from database", zap.Error(err))
		return err
	}
	logger.Log().Info("Disconnected from database ...")
	return nil
}
