package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/crud-gin/src/configuration/validation"
	"github.com/hugovallada/crud-gin/src/controller/model/request"
)

func GetUserById(c *gin.Context) {
}

func GetUserByEmail(c *gin.Context) {}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("error trying to marshal object, error=%s\n", err.Error())
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}

	c.JSON(201, userRequest)
}

func DeleteUser(c *gin.Context) {
}

func UpdateUser(c *gin.Context) {}
