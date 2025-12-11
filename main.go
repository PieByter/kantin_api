package main

import (
	"kantin_api/config"
	"kantin_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
