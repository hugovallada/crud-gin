package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/crud-gin/src/configuration/database/mongodb"
	"github.com/hugovallada/crud-gin/src/configuration/logger"
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
	// Init dependencies
	db, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Error("error trying o connect to database", err)
		return
	}

	userController := wire(db)

	router := gin.Default() // New não instancia handlers ou middlewares, o default instancia o middleware de logger e recovery
	routes.InitRoutes(&router.RouterGroup, userController)

	logger.Info("Starting application on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
