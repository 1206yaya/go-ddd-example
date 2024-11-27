package integration

import (
	"context"
	"testing"

	"github.com/1206yaya/go-ddd-example/internal/products/entities"
	"github.com/1206yaya/go-ddd-example/internal/products/repository"
	"github.com/1206yaya/go-ddd-example/internal/tests/integration/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductRepository_StoreProduct(t *testing.T) {
	// テストDBのセットアップ
	db, err := testutil.InitTestDB()
	require.NoError(t, err, "Failed to initialize test database")

	// テスト終了時のクリーンアップ
	t.Cleanup(func() {
		err := testutil.CleanupTestDB(db)
		if err != nil {
			t.Errorf("Failed to cleanup test database: %v", err)
		}
	})

	repo := repository.NewProductRepository(db)

	tests := []struct {
		name        string
		product     entities.Product
		wantErr     bool
		expectedErr error
	}{
		{
			name: "success_store_product",
			product: entities.Product{
				Name:  "Test Product",
				Price: 1000,
			},
			wantErr:     false,
			expectedErr: nil,
		},
		{
			name: "fail_store_product_empty_name",
			product: entities.Product{
				Name:  "",
				Price: 1000,
			},
			wantErr:     true,
			expectedErr: entities.ErrorProductNameEmpty,
		},
		{
			name: "fail_store_product_invalid_price",
			product: entities.Product{
				Name:  "Test Product",
				Price: 0,
			},
			wantErr:     true,
			expectedErr: entities.ErrorInvalidPrice,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 各テストケース前にデータをクリア
			err := testutil.TruncateTable(db, &entities.Product{})
			require.NoError(t, err, "Failed to truncate test database")

			// テストの実行
			err = repo.StoreProduct(context.Background(), tt.product)

			if tt.wantErr {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.expectedErr)
				return
			}

			assert.NoError(t, err)

			// 保存された製品の検証
			var storedProduct entities.Product
			err = db.First(&storedProduct, "name = ?", tt.product.Name).Error
			require.NoError(t, err, "Failed to retrieve stored product")

			assert.Equal(t, tt.product.Name, storedProduct.Name)
			assert.Equal(t, tt.product.Price, storedProduct.Price)
		})
	}
}
