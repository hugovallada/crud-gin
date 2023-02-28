package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
	"github.com/hugovallada/crud-gin/src/model"
	"github.com/hugovallada/crud-gin/src/model/repository/entity"
	"github.com/hugovallada/crud-gin/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (r *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init create user repository")
	collection := r.databaseConnection.Collection(os.Getenv(DATABASE_COLLECTION))

	userEntity := entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUser"), zap.String("email", email))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error("Internal error while connecting to the database", err, zap.String("journey", "findUser"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(userEntity), nil

}

func (r *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init create user repository")
	collection := r.databaseConnection.Collection(os.Getenv(DATABASE_COLLECTION))

	userEntity := entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUser"), zap.String("id", id))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error("Internal error while connecting to the database", err, zap.String("journey", "findUser"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	return converter.ConvertEntityToDomain(userEntity), nil

}
