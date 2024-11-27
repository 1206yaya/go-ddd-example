package repository

import (
	"context"

	"github.com/1206yaya/go-ddd-example/internal/products"
	"github.com/1206yaya/go-ddd-example/internal/products/entities"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) products.ProductRepository {
	return &repository{db: db}
}

func (r *repository) StoreProduct(ctx context.Context, product entities.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}

	result := r.db.WithContext(ctx).Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repository) GetProductByName(ctx context.Context, name string) (*entities.Product, error) {
	var product entities.Product
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
