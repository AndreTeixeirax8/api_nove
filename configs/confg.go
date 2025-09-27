package configs

import (
	"log"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DCPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JwtExperesIn  int //Tempo de validação do token
	TokenAuth     *jwtauth.JWTAuth
}

// Fnção para carregar as configurações
func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		log.Println("❌ Erro ao ler as configurações do viper.ReadConfig")
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Println("❌ Erro ao usar o unmarshal no cfg")
		panic(err)
	}

	//Cria uma instancia para poder gerar o token jwt
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err

}
