package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func GetDBConn(ctx context.Context, log *zap.SugaredLogger, dbURI string, dbName string) (*mongo.Database, error) {
	if dbURI == "" {
		return nil, errors.New("Set proper mongoDB URI")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, fmt.Errorf("Error occured connecting database: %v", err)
	}

	database := client.Database(dbName)
	return database, nil
}
