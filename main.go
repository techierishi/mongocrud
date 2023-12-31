package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/techierishi/mongocrud/http"
	"github.com/techierishi/mongocrud/repository"
)

var (
	DbPath  = getEnv("DBPATH", "mongodb://127.0.0.1:27017")
	SvcPort = getEnv("SVCPORT", "8060")
)

func main() {
	// create a database connection
	client, err := mongo.NewClient(options.Client().ApplyURI(DbPath))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	// create a repository
	repository := repository.NewRepository(client.Database("users"))

	// create an http server
	server := http.NewServer(repository)

	// create a gin router
	router := gin.Default()

	{
		router.GET("/users/:email", server.GetUser)
		router.GET("/users", server.GetUsers)
		router.POST("/users", server.CreateUser)
		router.PUT("/users/:email", server.UpdateUser)
		router.DELETE("/users/:email", server.DeleteUser)
	}
	// start the router
	router.Run(":" + SvcPort)
}
