package main

import (
	"dz/4-order-api/configs"
	"dz/4-order-api/internal/product"
	"dz/4-order-api/pkg/db"
	"dz/4-order-api/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)
	router := http.NewServeMux()

	//providers
	productRepo := product.NewProductRepository(db)

	//handlers
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepo,
	})

	server := http.Server{
		Addr:    config.Port,
		Handler: middleware.Logger(router),
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
