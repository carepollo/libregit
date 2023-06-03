package utils

import "github.com/joho/godotenv"

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("couldn't load environment variables file")
	}
}
