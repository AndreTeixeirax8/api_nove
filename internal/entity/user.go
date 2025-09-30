package entity

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
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
		ID:       "id",
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

// Faz a comparacao entre a senha informada e o hash do banco
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
