package service

import (
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
	"github.com/hugovallada/crud-gin/src/model"
	"github.com/hugovallada/crud-gin/src/model/repository"
)

type userDomainService struct {
	userReposiotry repository.UserRepository
}

func NewUserDomainService(userReposittory repository.UserRepository) UserDomainService {
	return &userDomainService{
		userReposiotry: userReposittory,
	}
}

type UserDomainService interface {
	Create(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
