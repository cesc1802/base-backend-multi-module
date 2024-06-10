package rest

import (
	"net/http"

	"github.com/cesc1802/auth-module/feature/authentication/dto"
	"github.com/cesc1802/auth-module/feature/authentication/storage"
	"github.com/cesc1802/auth-module/feature/authentication/usecase"
	"github.com/cesc1802/share-module/system"
	"github.com/gin-gonic/gin"
)

func Register(mono system.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := mono.DB()

		store := storage.NewSqlStorage(db)
		uc := usecase.NewRegisterUserUseCase(store)

		if err := uc.Register(c.Request.Context(), dto.RegisterRequest{}); err != nil {
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
