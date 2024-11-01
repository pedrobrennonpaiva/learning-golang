package config

import (
	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port             string `env:"PORT" envDefault:"8080"`
	ConnectionString string `env:"CONNECTION_STRING"`
	SecretKey        string `env:"SECRET_KEY"`
}

var cfg *config

func Parse() *config {

	cfg = &config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}

	return cfg
}

func GetConfig() *config {
	return cfg

}
