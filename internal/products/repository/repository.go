package repository

import (
	"context"

	"github.com/1206yaya/go-ddd-example/internal/products"
	"github.com/1206yaya/go-ddd-example/internal/products/entities"
)

type repository struct {
}

func NewProductRepository() products.ProductRepository {
	return &repository{}
}

func (r *repository) StoreProduct(ctx context.Context, request entities.Product) error {
	return nil
}
