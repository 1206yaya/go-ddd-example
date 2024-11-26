package http

import (
	"net/http"

	"github.com/1206yaya/go-ddd-example/internal/products"
	"github.com/1206yaya/go-ddd-example/internal/products/dtos"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	uc products.ProductUsecase
}

func NewProductHandler(uc products.ProductUsecase) products.ProductHandlers {
	return &handlers{uc: uc}
}

func (h *handlers) CreateProduct(c echo.Context) error {
	err := h.uc.CreateProduct(c.Request().Context(), dtos.Product{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}
