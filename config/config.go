package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	Env      string
	Port     string
	Protocol string
}

type Db struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

type Setting struct {
	Server Server
	Db     Db
}

func LoadConfig() *Setting {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	setting := Setting{
		Server: Server{
			Env:      os.Getenv("ENV"),
			Port:     os.Getenv("PORT"),
			Protocol: os.Getenv("PROTOCOL"),
		},
		Db: Db{
			DbHost: os.Getenv("DB_HOST"),
			DbPort: os.Getenv("DB_PORT"),
			DbUser: os.Getenv("DB_USER"),
			DbPass: os.Getenv("DB_PASS"),
			DbName: os.Getenv("DB_NAME"),
		},
	}
	return &setting

}
