package main

import (
	"github.com/hugovallada/crud-gin/src/controller"
	"github.com/hugovallada/crud-gin/src/model/repository"
	"github.com/hugovallada/crud-gin/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// Init dependencies

func wire(db *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(db)
	srv := service.NewUserDomainService(repo)
	return controller.NewUserController(srv)
}
