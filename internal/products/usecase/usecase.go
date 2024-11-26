package usecase

import (
	"context"

	"github.com/1206yaya/go-ddd-example/internal/products"
	"github.com/1206yaya/go-ddd-example/internal/products/dtos"
	"github.com/1206yaya/go-ddd-example/internal/products/entities"
)

type usecase struct {
	repo products.ProductRepository
}

func NewProductUsecase(r products.ProductRepository) products.ProductUsecase {
	return &usecase{repo: r}
}

func (uc *usecase) CreateProduct(ctx context.Context, request dtos.Product) error {

	return uc.repo.StoreProduct(ctx, entities.Product{})
}
