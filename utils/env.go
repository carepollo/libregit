package utils

import "github.com/joho/godotenv"

// use dotenv to load .env file declarations into environment variables
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("couldn't load environment variables file")
	}
}
