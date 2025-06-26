package main

import (
	"log/slog"
	"net/http"

	"github.com/LealKevin/simple-api/internal/category"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/ping", ping)
	category.Init(e)

	if err := e.Start(":8080"); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}

func ping(c echo.Context) error {
	c.String(http.StatusOK, "Pong")
	return nil
}
