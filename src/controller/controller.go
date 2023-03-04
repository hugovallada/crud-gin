package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/crud-gin/src/configuration/logger"
	"github.com/hugovallada/crud-gin/src/configuration/rest_err"
	"github.com/hugovallada/crud-gin/src/configuration/validation"
	"github.com/hugovallada/crud-gin/src/controller/model/request"
	"github.com/hugovallada/crud-gin/src/model"
	"github.com/hugovallada/crud-gin/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

var (
	UserDomainInterface *model.UserDomainInterface
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Finding user by id...", zap.String("journey", "findUserByID"))
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex("6403a2a6aca9f6476d99ab33"); err != nil {
		errorMessage := rest_err.NewBadRequestError("USERID is not valid")
		c.AbortWithStatusJSON(errorMessage.Code, errorMessage)
		logger.Error("Error trying to validate user", err, zap.Any("journey", "findUserById"))
		return
	}
	userDomain, err := uc.service.FindUserByID(userId)
	if err != nil {
		logger.Error("Error trying to find the user", err, zap.Any("journey", "findUserById"))
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Finding user by id...", zap.String("journey", "findUserByID"))
	email := c.Param("email")
	if _, err := mail.ParseAddress(email); err != nil {
		errorMessage := rest_err.NewBadRequestError("email is not valid")
		c.AbortWithStatusJSON(errorMessage.Code, errorMessage)
		logger.Error("Error trying to validate user", err, zap.Any("journey", "findUserByEmail"))
		return
	}
	userDomain, err := uc.service.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to find the user", err, zap.Any("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, userDomain)
}

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
	domainResult, err := uc.service.Create(domain)
	if err != nil {
		logger.Error("error trying to create an object", err, zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}
	response := view.ConvertDomainToResponse(domainResult)
	c.JSON(http.StatusCreated, response)
	logger.Info("User created", zap.String("email", response.Email))
}

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
}

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {

}
