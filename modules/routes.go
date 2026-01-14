package modules

import (
	"github.com/gin-gonic/gin"
	"glasdou.wtf/template/modules/health"
)

func RegisterRoutes(version *gin.RouterGroup) {
	health.RegisterRoutes(version)
}
