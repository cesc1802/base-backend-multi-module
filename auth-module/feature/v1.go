package feature

import (
	"net/http"

	registerRest "github.com/cesc1802/auth-module/feature/authentication/transport/rest"
	inspectRest "github.com/cesc1802/auth-module/feature/inspection/transport/rest"
	"github.com/cesc1802/share-module/system"
	"github.com/gin-gonic/gin"
)

func RegisterHandlerV1(mono system.Service) {
	router := mono.Router()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"data": "success",
		})
	})
	v1 := router.Group("/api/v1")
	authService := v1.Group("/auth-service")
	{
		authService.POST("/register", registerRest.Register(mono))
		authService.POST("/inspect", inspectRest.InspectToken(mono))
	}

}
