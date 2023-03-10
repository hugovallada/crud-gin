package repository

import (
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
	"github.com/hugovallada/crud-gin/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE_COLLECTION = "DATABASE_COLLECTION"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}
