package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, healthCheck())
}

func getDbConnection(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, checkDbConnection())
}
