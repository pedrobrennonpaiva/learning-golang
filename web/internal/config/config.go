package config

import (
	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	ApiUrl string `env:"API_URL"`
	Port   string `env:"PORT"`
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
