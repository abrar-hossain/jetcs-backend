package main

import (
	"jetcs-backend/config"
	"jetcs-backend/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	server.RegisterRoutes(r)
	r.Run(":8080")
}
