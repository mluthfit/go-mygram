package main

import (
	"go-mygram/controllers"
	"go-mygram/databases"
	"go-mygram/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const PORT = ":8000"

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	var db = databases.NewDatabase()
	var router = gin.Default()
	var server = controllers.NewServer(router, db)

	server.InitiateRoutes(routers.ApiRoutes)
	server.Run(PORT)
}
