package system

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"share-module/config"
	"share-module/waiter"
)

type Service interface {
	Config() config.AppConfig
	DB() *gorm.DB
	Router() *gin.Engine
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Service) error
}
