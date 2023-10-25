package config

import (
	"log"

	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload"
)

type DBEngine string

const (
	SQLite   DBEngine = "sqlite"
	Postgres DBEngine = "postgres"
)

type Config struct {
	Port         int      `env:"PORT" envDefault:"8080"`
	DBEngine     DBEngine `env:"DB_ENGINE" envDefault:"sqlite"`
	IsProduction bool     `env:"PRODUCTION" envDefault:"false"`
	Secret       string   `env:"SECRET" envDefault:"secret"`
	DBConnStr    string   `env:"DB_CONN_STRING" envDefault:"postgres://user:pass@localhost:5432/db?sslmode=disable"`
}

var Cfg Config

func ReadEnvVars() Config {

	Cfg = Config{}
	if err := env.Parse(&Cfg); err != nil {
		log.Fatal(err)
	}

	return Cfg
}
