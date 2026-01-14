package health

import (
	"context"
	"time"

	"glasdou.wtf/template/config"
	"glasdou.wtf/template/modules/common/types"
)

func healthCheck() types.HealthCheckResponse {
	return types.HealthCheckResponse{"status": "ok"}
}

func checkDbConnection() types.HealthCheckResponse {
	client := config.MongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return types.HealthCheckResponse{"database": "disconnected", "error": err.Error()}
	}

	return types.HealthCheckResponse{"database": "connected"}
}

// Test lefthook with staged files
