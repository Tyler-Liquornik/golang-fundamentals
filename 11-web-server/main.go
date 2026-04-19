package main

import (
	"github.com/Tyler-Liquornik/golang-fundamentals/11-web-server/db"
	"github.com/Tyler-Liquornik/golang-fundamentals/11-web-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
