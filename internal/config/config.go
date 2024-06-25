package config

import (
	"flag"
	"os"

	"github.com/AntonovIv/post_graphQlservice/internal/db/postgre"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string         `env:"ENV" type-defolt:"local"`
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
	var res string

	flag.StringVar(&res, "db", "", "what db to use: memory or postgres")
	flag.Parse()

	if res != "" {
		os.Setenv("DB_TYPE", res)
	}
}
