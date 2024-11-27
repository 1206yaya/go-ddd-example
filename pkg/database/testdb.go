package database

import (
	"fmt"

	"github.com/1206yaya/go-ddd-example/internal/products/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TableNamer interface {
	TableName() string
}

func InitTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open test database: %w", err)
	}

	// Auto-migrate for test database
	if err := db.AutoMigrate(&entities.Product{}); err != nil {
		return nil, fmt.Errorf("failed to migrate test database: %w", err)
	}

	return db, nil
}

// truncateTable は指定されたテーブルのデータを削除します
func truncateTable(tx *gorm.DB, tableName string) error {
	// 外部キー制約を一時的に無効化
	if err := tx.Exec("PRAGMA foreign_keys = OFF").Error; err != nil {
		return fmt.Errorf("failed to disable foreign keys: %w", err)
	}
	defer func() {
		// 関数終了時に外部キー制約を再度有効化
		if err := tx.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
			// deferredエラーはログに記録するべきですが、
			// この例では簡略化のため省略しています
		}
	}()

	// テーブルのデータを削除
	if err := tx.Exec(fmt.Sprintf("DELETE FROM %s", tableName)).Error; err != nil {
		return fmt.Errorf("failed to truncate table %s: %w", tableName, err)
	}

	// Auto Incrementをリセット
	if err := tx.Exec(fmt.Sprintf("DELETE FROM sqlite_sequence WHERE name='%s'", tableName)).Error; err != nil {
		return fmt.Errorf("failed to reset auto increment for table %s: %w", tableName, err)
	}

	return nil
}

// TruncateAllTables はデータベースの全テーブルをクリアします
func TruncateAllTables(db *gorm.DB) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return fmt.Errorf("failed to get tables: %w", err)
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, table := range tables {
			if table == "sqlite_sequence" {
				continue
			}
			if err := truncateTable(tx, table); err != nil {
				return err
			}
		}
		return nil
	})
}

// TruncateTable は指定されたテーブルをクリアします
func TruncateTable(db *gorm.DB, model TableNamer) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return truncateTable(tx, model.TableName())
	})
}

// CleanupTestDB はテストデータベースをクリーンアップします
func CleanupTestDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close test database connection: %w", err)
	}

	return nil
}
