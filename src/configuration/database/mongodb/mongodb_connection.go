package mongodb

import (
	"context"

	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection() {
	ctx := context.Background()
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic("Banco de dados não configurado")
	}
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	logger.Info("Conexão com o banco feita com sucesso")
}
