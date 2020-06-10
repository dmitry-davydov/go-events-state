package api

import (
	"context"
	"fmt"

	v1 "github.com/dmitry-davydov/go-events-state/internal/api/v1"
	"github.com/dmitry-davydov/go-events-state/internal/storage"
	"github.com/labstack/echo/v4"
)

// CreateServer создает обтработчик http запросов
func CreateServer(port int, context context.Context) {
	e := echo.New()

	v1Group := e.Group("/api/v1")

	store := storage.NewStorage(5, context)

	v1.New(v1Group, store)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
