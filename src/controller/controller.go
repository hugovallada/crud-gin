package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"github.com/hugovallada/crud-gin/src/configuration/validation"
	"github.com/hugovallada/crud-gin/src/controller/model/request"
	"github.com/hugovallada/crud-gin/src/model"
	"github.com/hugovallada/crud-gin/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface *model.UserDomainInterface
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {}

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Creating user...", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to marshal object", err, zap.String("journey", "createUser"))
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}
	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	if err := uc.service.Create(domain); err != nil {
		logger.Error("error trying to create an object", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}
	response := view.ConvertDomainToResponse(domain)
	c.JSON(http.StatusCreated, response)
	logger.Info("User created", zap.String("email", response.Email))
}

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
}

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {}
