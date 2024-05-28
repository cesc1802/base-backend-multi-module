package rest

import (
	"net/http"

	"auth-module/feature/authentication/dto"
	"auth-module/feature/authentication/storage"
	"auth-module/feature/authentication/usecase"
	"github.com/gin-gonic/gin"
	"share-module/system"
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
