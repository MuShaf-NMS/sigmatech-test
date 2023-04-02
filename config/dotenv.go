package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error .env file")
	}
}

func getVariable(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("Environtment variable not found")
}
