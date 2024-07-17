package configs

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type PgConfig struct {
	Host     string `env:"PG_HOST"`
	Port     string `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Dbname   string `env:"PG_DBNAME"`
}
type Config struct {
	Pg PgConfig
}

func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	config := new(Config)
	err = env.Parse(config)
	if err != nil {
		panic(err)
	}

	return config
}
