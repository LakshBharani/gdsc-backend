package main

import (
	"github.com/gin-gonic/gin"
	"laksh.com/db"
	"laksh.com/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
