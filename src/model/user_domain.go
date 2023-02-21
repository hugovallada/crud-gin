package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
	"go.uber.org/zap"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	EncryptPassword()
}

type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) Create() *rest_err.RestErr {
	logger.Info("init create user model", zap.String("journey", "createUser"))
	u.EncryptPassword()
	logger.Info(u.Password)
	return nil
}

func (u *userDomain) DeleteUser(string) *rest_err.RestErr {
	return nil
}

func (u *userDomain) FindUser(string) (*userDomain, *rest_err.RestErr) {
	return nil, nil
}

func (u *userDomain) UpdateUser(string) *rest_err.RestErr {
	return nil
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}
