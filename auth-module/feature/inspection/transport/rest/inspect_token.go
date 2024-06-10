package rest

import (
	"net/http"

	"github.com/cesc1802/auth-module/feature/inspection/storage"
	"github.com/cesc1802/auth-module/feature/inspection/usecase"
	"github.com/cesc1802/share-module/system"
	"github.com/gin-gonic/gin"
)

func InspectToken(mono system.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := mono.DB()
		tokProvider := mono.TokenProvider()

		store := storage.NewSqlStorage(db)
		uc := usecase.NewInspectToken(store, tokProvider)

		if err := uc.Inspect(c.Request.Context(), "dto.RegisterRequest{}"); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{
				"message": "bad request",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"message": "success",
		})
		return
	}
}
