package usecase

import (
	"context"

	"github.com/1206yaya/go-ddd-example/internal/products"
	"github.com/1206yaya/go-ddd-example/internal/products/dtos"
	"github.com/1206yaya/go-ddd-example/internal/products/mapper"
)

type usecase struct {
	repo   products.ProductRepository
	mapper mapper.ProductMapper
}

func NewProductUsecase(r products.ProductRepository, m mapper.ProductMapper) products.ProductUsecase {
	return &usecase{repo: r, mapper: m}
}

func (uc *usecase) CreateProduct(ctx context.Context, request dtos.CreateProductRequest) error {
	product := uc.mapper.ToEntity(request)
	return uc.repo.StoreProduct(ctx, product)
}
