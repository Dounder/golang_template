package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"glasdou.wtf/template/config"
	"glasdou.wtf/template/modules"
)

func main() {
	config.MustLoad()
	log.Println("Configuration loaded successfully")

	if err := config.LoadConfig(); err != nil {
		log.Fatalf("App configuration validation failed:\n\n%v", err)
	}

	config.InitDatabase()
	log.Println("Database initialized successfully")

	gin.SetMode(config.Envs.App.GinMode)
	server := gin.Default()
	v1 := server.Group("/api/v1")

	modules.RegisterRoutes(v1)

	port := config.Envs.Server.Port
	server.Run(fmt.Sprintf(":%d", port))
	log.Printf("Server running on port %d", port)
}
