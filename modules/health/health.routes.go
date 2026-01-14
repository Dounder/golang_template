package health

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(version *gin.RouterGroup) {
	healthGroup := version.Group("/health")

	healthGroup.GET("", getHealth)
	healthGroup.GET("/db", getDbConnection)
}
