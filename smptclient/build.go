package smtpclient

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// createSMTPConfig loads SMTP configuration details from environment variables
// and returns a populated smptConfig struct.
func createSMTPConfig() (smptConfig, error) {

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

// createPlainAuthentication creates an SMTP Plain authentication
// using the provided smptConfig and returns an smtp.Auth instance.
func createPlainAuthentication(client smptConfig) smtp.Auth {
	return smtp.PlainAuth(
		client.identity,
		client.username,
		client.password,
		client.Host,
	)
}

// BuildClient loads environment variables and creates a configured SMTP client.
// It returns the client and any errors encountered during the process.
func BuildClient(envPath string) (Client, error) {

	if err := godotenv.Load(envPath); err != nil {
		return &smptClient{}, fmt.Errorf("an error has occured while loading env variables: %w", err)
	}

	config, err := createSMTPConfig()

	if err != nil {
		return &smptClient{}, fmt.Errorf("BuildClient: %w", err)
	}

	auth := createPlainAuthentication(config)

	return &smptClient{
		config: config,
		auth:   auth,
	}, nil

}

// BuildMessage constructs an email message with the given recipients, subject, and body.
// It returns the message as a slice of bytes for sending it via SMTP.
func BuildMessage(recipients []string, subject, body string) []byte {

	CRLF := "\r\n"

	message := []string{
		fmt.Sprintf("To: %s", strings.Join(recipients, ",")),
		fmt.Sprintf("Subject: %s", subject),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		body,
	}

	return []byte(strings.Join(message, CRLF))

}
