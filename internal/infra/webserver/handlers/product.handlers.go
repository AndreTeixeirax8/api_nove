package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/api_nove/internal/dto"
	"github.com/api_nove/internal/entity"
	database "github.com/api_nove/internal/infra/db"
	entityPkg "github.com/api_nove/pkg/entity"
	"github.com/go-chi/chi/v5"
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

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") //Obtem parametro da rota
	if id == "" {
		log.Println("⚠️[HANDLER] Id informado esta vazio")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		log.Println("🔍[HANDLER] Não foi encontrado um produto com o id informado")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product) //Retorna o JSON do produto
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		log.Println("⚠️[HANDLER] Id informado esta vazio")
		w.WriteHeader((http.StatusBadRequest))
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		log.Println("⚠️[HANDLER] Id informado esta vazio")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println("🔍[HANDLER] Não foi encontrado um produto com o id informado")
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("❌[HANDLER] Ocorreu um erro ao deletar o produto")
		return
	}

	w.WriteHeader(http.StatusOK)
}
