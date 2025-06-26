package category

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/categories", h.GetAll)
	e.GET("/category/:id", h.GetByID)
	e.POST("/category", h.Create)
	e.DELETE("/category/:id", h.Delete)
}

func (h *Handler) GetAll(c echo.Context) error {
	categories, err := h.Service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not load categories")
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *Handler) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	category, err := h.Service.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, category)
}

func (h *Handler) Create(c echo.Context) error {
	var category Category
	err := c.Bind(&category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := h.Service.Create(category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "category created successfuly",
	})
}

func (h *Handler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	categories, err := h.Service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.Service.Delete(id, categories); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "category deleted",
	})
}
