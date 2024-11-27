package main

import (
	"github.com/1206yaya/go-ddd-example/internal/products/delivery/http"
	"github.com/1206yaya/go-ddd-example/internal/products/mapper"
	"github.com/1206yaya/go-ddd-example/internal/products/repository"
	"github.com/1206yaya/go-ddd-example/internal/products/usecase"
	"github.com/1206yaya/go-ddd-example/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	dbConfig := database.Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "password",
		DBName:   "product_db",
	}
	// データベース接続
	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	echo := echo.New()

	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUsecase(productRepo, mapper.NewProductMapper())
	productHandlers := http.NewProductHandler(productUC)

	echo.POST("/product", productHandlers.CreateProduct)

	log.Info(echo.Start(":8888"))
}
