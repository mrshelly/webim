package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck return status of chat server
func HealthCheck(c *gin.Context, appService *ServiceProvider) {
	c.JSON(http.StatusOK, CommonResponse{
		Message: InfoHealthCheck,
	})
}
