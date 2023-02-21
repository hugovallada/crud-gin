package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hugovallada/crud-gin/src/model/service"
)

type UserControllerInterface interface {
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

func NewUserController(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}
