package core

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/plaid/go-envvar/envvar"
)

type Env struct {
	HttpHost string `envvar:"HTTP_HOST"`
}

func GetEnv() *Env {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("can't open .env file, using system environment variables")
	}

	env := &Env{}
	if err := envvar.Parse(env); err != nil {
		log.Fatal("can't parse environment variables")
	}

	return env
}
