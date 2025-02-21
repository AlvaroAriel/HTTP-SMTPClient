package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Client interface {
	SendEmail(recipients []string, message []byte) error
}

type smptClient struct {
	config smptConfig
	auth   smtp.Auth
}

func (c smptClient) SendEmail(recipients []string, message []byte) error {

	err := smtp.SendMail(
		c.config.Address,
		c.auth,
		c.config.username,
		recipients,
		message,
	)

	if err != nil {
		return fmt.Errorf("an error has occured while sending email %w", err)
	}

	return nil
}

type smptConfig struct {
	Host     string
	Address  string
	identity string
	username string
	password string
}

type Email struct {
	Subject string
	Body    string
}

func (e Email) BuildMessage(recipients []string) []byte {

	CRLF := "\r\n"

	message := []string{
		fmt.Sprintf("To: %s", strings.Join(recipients, ",")),
		fmt.Sprintf("Subject: %s", e.Subject),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		e.Body,
	}

	return []byte(strings.Join(message, CRLF))

}
func CreateSMTPConfig(filename string) (smptConfig, error) {

	if err := godotenv.Load(); err != nil {
		return smptConfig{}, fmt.Errorf("an error has occured while loading env variables: %w", err)
	}

	host := os.Getenv("SMTPC_SERVER")
	port := os.Getenv("SMTPC_PORT")
	identity := os.Getenv("SMTPC_IDENTITY")
	username := os.Getenv("SMTPC_USERNAME")
	password := os.Getenv("SMTPC_PASSWORD")

	addr := strings.Join([]string{host, port}, ":")

	return smptConfig{
		Host:     host,
		Address:  addr,
		identity: identity,
		username: username,
		password: password,
	}, nil
}
func CreatePlainAuthentication(client smptConfig) smtp.Auth {
	return smtp.PlainAuth(
		client.identity,
		client.username,
		client.password,
		client.Host,
	)
}

func BuildClient(envPath string) (Client, error) {
	config, err := CreateSMTPConfig(envPath)

	if err != nil {
		return smptClient{}, fmt.Errorf("BuildClient: %w", err)
	}

	auth := CreatePlainAuthentication(config)

	return smptClient{
		config: config,
		auth:   auth,
	}, nil

}

func main() {
	envPath := ".env"
	recipient := []string{"some emails :)"}

	client, err := BuildClient(envPath)

	if err != nil {
		log.Fatal(err)
	}
	e := Email{
		Subject: "Subject",
		Body:    "message",
	}
	message := e.BuildMessage(recipient)
	err = client.SendEmail(recipient, message)

	if err != nil {
		fmt.Println("Failed sending email: ", err)
	}

}
