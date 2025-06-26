package category

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo) {
	store := NewMemoryStore()
	service := NewService(store)
	handler := NewHandler(service)
	handler.RegisterRoutes(e)
}
