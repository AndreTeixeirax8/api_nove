package entity

import (
	"errors"
	"time"

	"github.com/api_nove/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("ID é requerido")
	ErrInvalidID       = errors.New("ID é invalido")
	ErrNameIsRequired  = errors.New("Nome é obrigatorio")
	ErrPriceIsRequerid = errors.New("Preço é obrigatorio")
	ErrInvalidPrice    = errors.New("Preço invalido")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewIdD(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequerid
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
