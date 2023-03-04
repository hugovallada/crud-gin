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

func (u *userDomainService) FindUserByID(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init find user by id", zap.String("journey", "findUserById"))
	return u.userReposiotry.FindUserByID(userId)
}

func (u *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init find user by email", zap.String("journey", "findUserByEmail"))
	return u.userReposiotry.FindUserByEmail(email)
}

func (u *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	return nil
}
