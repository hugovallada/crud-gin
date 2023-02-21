package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugovallada/crud-gin/src/controller"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.GET("/users/id/:userId", controller.GetUserById)
	rg.GET("/users/email/:userEmail", controller.GetUserByEmail)
	rg.POST("/users", controller.CreateUser)
	rg.PUT("/users/id/:userid", controller.UpdateUser)
	rg.DELETE("/users/id/:userid", controller.DeleteUser)
}
