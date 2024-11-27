package mapper

import (
	"github.com/1206yaya/go-ddd-example/internal/products/dtos"
	"github.com/1206yaya/go-ddd-example/internal/products/entities"
)

type ProductMapper interface {
	ToEntity(dto dtos.CreateProductRequest) entities.Product
	ToDTO(entity entities.Product) dtos.CreateProductResponse
}

type productMapper struct{}

func NewProductMapper() ProductMapper {
	return &productMapper{}
}

func (m *productMapper) ToEntity(dto dtos.CreateProductRequest) entities.Product {
	return entities.Product{
		Name:  dto.Name,
		Price: dto.Price,
	}
}

func (m *productMapper) ToDTO(entity entities.Product) dtos.CreateProductResponse {
	return dtos.CreateProductResponse{
		ID:    entity.ID,
		Name:  entity.Name,
		Price: entity.Price,
	}
}
