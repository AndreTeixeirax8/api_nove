package main

import (
	"log"

	"github.com/api_nove/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	log.Println(config.DBDriver)
}
