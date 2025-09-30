package entity

import (
	"log"

	"github.com/api_nove/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"` //Usa o pacote publico para tornar a funcao new id publica
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

// Retorna um usuario e um nil ou um nil e um erro (No caso nil é = a vazio)
func NewUser(name, email, password string) (*User, error) {
	//Pega o password e converte em um hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("❌ Erro ao converter a senha em hash")
		return nil, err
	}

	return &User{
		ID:       entity.NewIdD(), //Gera um novo id
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

// Faz a comparacao entre a senha informada e o hash do banco retorna true ou false
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
