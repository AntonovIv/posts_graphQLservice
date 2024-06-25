package config

import "github.com/joho/godotenv"

func mustinitCfg(path string) {
	if err := godotenv.Load(path); err != nil {
		panic("cant init env" + err.Error())
	}

}
