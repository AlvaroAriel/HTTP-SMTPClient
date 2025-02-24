package config

import (
	"log"
	"net"
	"os"
)

type Config struct {
	Address    string
	Enviroment string
}

func NewConfig() *Config {

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	env := os.Getenv("APP_ENV")

	if host == "" {
		log.Fatal("no host given")
	}

	return &Config{
		Address:    net.JoinHostPort(host, port),
		Enviroment: env,
	}
}
