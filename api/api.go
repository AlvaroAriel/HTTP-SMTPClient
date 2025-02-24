package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlvaroAriel/HTTP-SMTPClient/config"
	"github.com/joho/godotenv"

	"github.com/AlvaroAriel/HTTP-SMTPClient/internal/email"
	"github.com/AlvaroAriel/HTTP-SMTPClient/internal/middleware"

	smtpclient "github.com/AlvaroAriel/HTTP-SMTPClient/smptclient"
)

func newServer(config *config.Config, smtpClient smtpclient.Client) http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /send-email", email.HandleSendEmail(smtpClient))

	handler := middleware.CorsMiddleware(mux)

	return http.Server{
		Addr:    config.Address,
		Handler: handler,
	}

}

func Run(envPaths ...string) {

	if err := godotenv.Load(envPaths...); err != nil {
		log.Fatal("enviroment file not found")
	}

	config := config.NewConfig()
	smtpClient, err := smtpclient.BuildClient()

	if err != nil {
		log.Fatal("smtp client build failed")
	}

	server := newServer(config, smtpClient)

	fmt.Printf("Server running on address: %s in %s enviroment\n", config.Address, config.Enviroment)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
