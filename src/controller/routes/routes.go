package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugovallada/crud-gin/src/controller"
)

func InitRoutes(rg *gin.RouterGroup, controller controller.UserControllerInterface) {
	rg.GET("/users/id/:userId", controller.FindUserById)
	rg.GET("/users/email/:userEmail", controller.FindUserByEmail)
	rg.POST("/users", controller.CreateUser)
	rg.PUT("/users/id/:userid", controller.UpdateUser)
	rg.DELETE("/users/id/:userid", controller.DeleteUser)
}
