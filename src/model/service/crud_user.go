package service

import (
	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
	"github.com/hugovallada/crud-gin/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) Create(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init create user model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	response, err := u.userReposiotry.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call the user repository", err, zap.String("journey", "createUser"))
		return nil, err
	}
	return response, nil
}

func (u *userDomainService) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}

func (u *userDomainService) FindUser(userId string) (*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (u *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	return nil
}
