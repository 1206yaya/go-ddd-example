package integration

import (
	"context"
	"testing"

	"github.com/1206yaya/go-ddd-example/internal/products/dtos"
	"github.com/1206yaya/go-ddd-example/internal/products/entities"
	"github.com/1206yaya/go-ddd-example/internal/products/mapper"
	"github.com/1206yaya/go-ddd-example/internal/products/repository"
	"github.com/1206yaya/go-ddd-example/internal/products/usecase"
	"github.com/1206yaya/go-ddd-example/internal/tests/integration/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
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

	// 依存関係の構築
	productRepo := repository.NewProductRepository(db)
	productMapper := mapper.NewProductMapper()
	productUC := usecase.NewProductUsecase(productRepo, productMapper)

	tests := []struct {
		name        string
		request     dtos.CreateProductRequest
		wantErr     bool
		expectedErr error
	}{
		{
			name: "正常系：製品の作成に成功",
			request: dtos.CreateProductRequest{
				Name:  "Test Product",
				Price: 1000,
			},
			wantErr: false,
		},
		{
			name: "異常系：製品名が空",
			request: dtos.CreateProductRequest{
				Name:  "",
				Price: 1000,
			},
			wantErr:     true,
			expectedErr: entities.ErrorProductNameEmpty,
		},
		{
			name: "異常系：価格が無効",
			request: dtos.CreateProductRequest{
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
			err = productUC.CreateProduct(context.Background(), tt.request)

			// エラー検証
			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					assert.ErrorIs(t, err, tt.expectedErr)
				}
				return
			}

			// 成功ケースの検証
			assert.NoError(t, err)

			// データベースに保存された製品の検証
			product, err := productRepo.GetProductByName(context.Background(), tt.request.Name)
			require.NoError(t, err, "Failed to retrieve stored product")

			assert.Equal(t, tt.request.Name, product.Name)
			assert.Equal(t, tt.request.Price, product.Price)
			assert.NotZero(t, product.ID)
			assert.NotZero(t, product.CreatedAt)
			assert.NotZero(t, product.UpdatedAt)
		})
	}
}
