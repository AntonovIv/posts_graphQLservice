package config

import (
	"flag"
	"os"

	"github.com/AntonovIv/post_graphQlservice/internal/db/postgre"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string         `env:"ENV" type-defolt:"local"`
	EnvFile  string         `env:"ENV_FILE" type-defolt:".env"`
	Port     int            `env:"PORT_APP" type-defolt:"8080"`
	DbConfig postgre.Config `env:"DB_CONFIG"`
}

func MustLoadCfg() *Config {
	mustinitCfg(".env")

	fetchDbType()

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("cant read config file" + err.Error())
	}
	return &cfg
}
func fetchDbType() {
	var db string
	var cont string

	flag.StringVar(&db, "db", "", "what db to use: memory or postgres")
	flag.StringVar(&cont, "container", "", "is it run in docker container?: y[yes]/n[no]")
	flag.Parse()

	if db != "" {
		os.Setenv("DB_TYPE", db)
	}
	if cont == "n" {
		os.Setenv("DB_HOSTNAME", "localhost")
	}
}
