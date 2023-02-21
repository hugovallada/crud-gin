package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/crud-gin/src/controller/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default() // New não instancia handlers ou middlewares, o default instancia o middleware de logger e recovery
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
