package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api_nove/internal/dto"
	"github.com/api_nove/internal/entity"
	database "github.com/api_nove/internal/infra/db"
)

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
