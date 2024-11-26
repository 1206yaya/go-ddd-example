package products

import (
	"context"

	"github.com/1206yaya/go-ddd-example/internal/products/entities"
)

type ProductRepository interface {
	StoreProduct(ctx context.Context, request entities.Product) error
}
