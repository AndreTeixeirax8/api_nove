package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api_nove/configs"
	"github.com/api_nove/internal/dto"
	"github.com/api_nove/internal/entity"
	database "github.com/api_nove/internal/infra/db"

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
	productHandler := NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)
	fmt.Printf("✅ Servidor rodando na porta %s\n", config.WebServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", config.WebServerPort), nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "❌ Dados do produto invalido", http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, "❌ Erro ao criar o produto", http.StatusInternalServerError)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		http.Error(w, "❌ Erro ao salvar o produto", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "✅ Produto criado com sucesso: %s", p.ID)
}
