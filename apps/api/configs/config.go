package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Env         string
	Addr        string
	DataBaseURL string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return &Config{
		AppName:     os.Getenv("APP_NAME"),
		Env:         os.Getenv("APP_ENV"),
		Addr:        os.Getenv("APP_ADDR"),
		DataBaseURL: os.Getenv("DATABASE_URL"),
	}
}
