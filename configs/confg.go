package configs

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
}
