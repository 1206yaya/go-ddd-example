package products

import (
	"context"

	"github.com/1206yaya/go-ddd-example/internal/products/dtos"
)

type ProductUsecase interface {
	CreateProduct(context.Context, dtos.CreateProductRequest) error
}
