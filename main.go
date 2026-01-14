package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"glasdou.wtf/template/config"
	"glasdou.wtf/template/modules"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	config.MustLoad()
	logger.Info("Configuration loaded successfully")

	if err := config.LoadConfig(); err != nil {
		logger.Error("App configuration validation failed:\n\n", "error", err)
	}

	config.InitDatabase()
	logger.Info("Database initialized successfully")
	gin.SetMode(config.Envs.App.GinMode)
	server := gin.Default()
	v1 := server.Group("/api/v1")

	modules.RegisterRoutes(v1)

	port := config.Envs.Server.Port
	if err := server.Run(fmt.Sprintf(":%d", port)); err != nil {
		logger.Error("Failed to start server:", "error", err)
	}

	logger.Info("Server is running on port", "port", port)
}
