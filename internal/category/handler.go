package category

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *Service
	Store   Store
}

func NewHandler(service *Service, store Store) *Handler {
	return &Handler{
		Service: service,
		Store:   store,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/categories", h.GetAll)
	e.GET("/category/:id", h.GetByID)
	e.POST("/category", h.HandleCreateRequest)
	e.DELETE("/category/:id", h.Delete)
}

func (h *Handler) GetAll(c echo.Context) error {
	categories, err := h.Store.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not load categories")
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *Handler) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	category, err := h.Store.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, category)
}

func (h *Handler) HandleCreateRequest(c echo.Context) error {
	var createCategoryRequest Category
	err := c.Bind(&createCategoryRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	createdCategory, err := h.Service.Create(createCategoryRequest.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.Store.Save(createdCategory); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "category created successfuly",
	})
}

func (h *Handler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.Store.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "category deleted",
	})
}
