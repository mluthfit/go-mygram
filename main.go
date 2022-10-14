package main

import (
	"go-mygram/controllers"
	"go-mygram/database"
	"go-mygram/routers"

	"github.com/gin-gonic/gin"
)

const PORT = ":8000"

func main() {
	var db = database.NewDatabase()
	var router = gin.Default()
	var server = controllers.NewServer(router, db)

	server.InitiateRoutes(routers.ApiRoutes)
	server.Run(PORT)
}
