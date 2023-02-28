package repository

import (
	"context"
	"os"

	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
	"github.com/hugovallada/crud-gin/src/model"
	"github.com/hugovallada/crud-gin/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (r *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init create user repository")
	collection := r.databaseConnection.Collection(os.Getenv(DATABASE_COLLECTION))
	value := converter.FromUserDomainToUserEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())

	}
	value.ID = result.InsertedID.(primitive.ObjectID)
	return converter.ConvertEntityToDomain(*value), nil
}
