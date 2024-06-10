package system

import (
	"context"

	"github.com/cesc1802/share-module/config"
	"github.com/cesc1802/share-module/tokprovider"
	"github.com/cesc1802/share-module/waiter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service interface {
	Config() config.AppConfig
	DB() *gorm.DB
	Router() *gin.Engine
	Waiter() waiter.Waiter
	TokenProvider() tokprovider.TokenProvider
}

type Module interface {
	Startup(context.Context, Service) error
}
