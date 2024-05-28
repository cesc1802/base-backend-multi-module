package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"net/http"
	"share-module/config"
	"share-module/waiter"
)

type System struct {
	cfg    config.AppConfig
	db     *gorm.DB
	router *gin.Engine
	waiter waiter.Waiter
}

func NewSystem(cfg config.AppConfig) *System {
	sys := &System{cfg: cfg}

	sys.initWaiter()
	sys.initMux()

	return sys
}

func (s *System) initWaiter() {
	s.waiter = waiter.New(waiter.CatchSignals())
}

func (s *System) initMux() {
	s.router = gin.New()
}

func (s *System) Config() config.AppConfig {
	return s.cfg
}

func (s *System) DB() *gorm.DB {
	return s.db
}

func (s *System) Router() *gin.Engine {
	return s.router
}

func (s *System) Waiter() waiter.Waiter {
	return s.waiter
}

func (s *System) WaitForWeb(ctx context.Context) error {
	webServer := &http.Server{
		Addr:    s.cfg.Web.Address(),
		Handler: s.router,
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Printf("web server started; listening at http://localhost%s\n", s.cfg.Web.Port)
		defer fmt.Println("web server shutdown")
		if err := webServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("web server to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
		defer cancel()
		if err := webServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return group.Wait()
}
