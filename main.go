package main

import (
	"example/event-management/db"
	"example/event-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
