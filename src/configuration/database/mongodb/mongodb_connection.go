package mongodb

import (
	"context"
	"os"

	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL = "DATABASE_URL"
	MONGO_DATABASE_NAME = "DATABASE_NAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbUri := os.Getenv(MONGODB_URL)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	logger.Info("Conexão com o banco feita com sucesso")
	return client.Database(os.Getenv(MONGO_DATABASE_NAME)), nil
}
