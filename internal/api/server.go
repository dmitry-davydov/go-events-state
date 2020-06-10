package api

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/dmitry-davydov/go-events-state/internal/api/v1"
	"github.com/dmitry-davydov/go-events-state/internal/models"
	"github.com/dmitry-davydov/go-events-state/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CreateServer создает обтработчик http запросов
func CreateServer(cfg models.Configuration) {
	e := echo.New()

	e.Use(middleware.Recover())

	if cfg.LogHttp {
		e.Use(middleware.Logger())
	}

	v1Group := e.Group("/api/v1")

	store := storage.NewStorage(cfg.InterruptedValidPeriod, cfg.DieAppContext)

	v1.AttachRoutes(v1Group, store)

	go func (){
		if err := e.Start(fmt.Sprintf(":%d", cfg.HttpServerPort)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	<- cfg.DieAppContext.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
