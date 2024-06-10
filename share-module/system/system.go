package system

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/cesc1802/share-module/config"
	"github.com/cesc1802/share-module/tokprovider"
	"github.com/cesc1802/share-module/waiter"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type System struct {
	cfg         config.AppConfig
	db          *gorm.DB
	router      *gin.Engine
	waiter      waiter.Waiter
	tokProvider tokprovider.TokenProvider
}

func newSystem(cfg config.AppConfig) *System {
	sys := &System{cfg: cfg}

	sys.initWaiter()
	sys.initGin()

	if err := sys.initDB(); err != nil {
		log.Fatalln("database cannot start: ", err)
	}

	return sys
}

func newSystemForMigrate(cfg config.AppConfig) *System {
	sys := &System{cfg: cfg}

	if err := sys.initDB(); err != nil {
		log.Fatalln("database cannot start: ", err)
	}

	return sys
}

func New(cfg config.AppConfig, action string) *System {
	if action == "migrate" {
		return newSystemForMigrate(cfg)
	}
	return newSystem(cfg)
}

func (s *System) initWaiter() {
	s.waiter = waiter.New(waiter.CatchSignals())
}

func (s *System) initGin() {
	if s.cfg.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	s.router = gin.New()
}

func (s *System) initDB() error {
	db, err := gorm.Open(postgres.Open(s.cfg.DB.Uri()), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return err
	}

	rawDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if err := rawDB.Ping(); err != nil {
		log.Fatalln(err)
		return err
	}

	s.db = db
	return nil
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
		fmt.Printf("web server started; listening at http://localhost:%s\n", s.cfg.Web.Port)
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

func (s *System) MigrateDB(fs fs.FS) error {
	goose.SetBaseFS(fs)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	rawDB, _ := s.db.DB()
	if err := goose.Up(rawDB, "."); err != nil {
		return err
	}

	return nil
}

func (s *System) TokenProvider() tokprovider.TokenProvider {
	return s.tokProvider
}
