package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
	"go.uber.org/zap"
)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetAge() int8 {
	return u.age
}

func (u *userDomain) GetID() string {
	return u.id
}

func (u *userDomain) Create() *rest_err.RestErr {
	logger.Info("init create user model", zap.String("journey", "createUser"))
	u.EncryptPassword()
	logger.Info(u.password)
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
	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}

func (u *userDomain) SetId(id string) {
	u.id = id
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}
