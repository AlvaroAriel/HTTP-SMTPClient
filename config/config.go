package config

import (
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

	return &Config{
		Address:    net.JoinHostPort(host, port),
		Enviroment: env,
	}
}
