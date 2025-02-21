package smtpclient

import (
	"fmt"
	"net/smtp"
)

// Client defines the methods for an SMTP client that can send emails.
type Client interface {
	SendEmail(recipients []string, message []byte) error
}

// smptClient is a concrete implementation of the Client interface
// with a configuration and authentication data for sending emails.
type smptClient struct {
	config smptConfig
	auth   smtp.Auth
}

// SendEmail sends an email to the specified recipients using the SMTP client configuration.
// It returns an error if the email fails to send.
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

// smptConfig holds the configuration data needed for an SMTP client.
// It includes the host, address, authentication details, and user credentials.
type smptConfig struct {
	Host     string
	Address  string
	identity string
	username string
	password string
}
