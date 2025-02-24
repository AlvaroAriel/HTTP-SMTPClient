package email

import (
	"log"
	"os"
	"strings"
)

func getRecipient() []string {
	recipients := os.Getenv("SMTPC_RECIPIENT")

	if recipients == "" {
		log.Fatal("no recipients found")
	}

	return strings.Split(recipients, ",")
}
