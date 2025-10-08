package main

import (
	"fmt"
	"net/http"

	"github.com/api_nove/configs"
	"github.com/api_nove/internal/entity"
	database "github.com/api_nove/internal/infra/db"
	"github.com/api_nove/internal/infra/webserver/handlers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)
	fmt.Printf("✅ Servidor rodando na porta %s\n", config.WebServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", config.WebServerPort), nil)
}
