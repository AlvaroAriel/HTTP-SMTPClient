package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	smtpclient "github.com/AlvaroAriel/HTTP-SMTPClient/smptclient"
)

type Email struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func getRecipient() ([]string, error) {
	recipients := os.Getenv("SMTPC_RECIPIENT")

	if recipients == "" {
		return nil, fmt.Errorf("no recipient")
	}

	return strings.Split(recipients, ","), nil
}

func main() {

	mux := http.NewServeMux()
	smtpClient, err := smtpclient.BuildClient(".env")

	if err != nil {
		fmt.Println(fmt.Errorf("error building smtp client"))
	}

	recipients, err := getRecipient()

	if err != nil {
		fmt.Println(fmt.Errorf("error getting recipient"))
	}

	mux.HandleFunc("POST /send", handleSendEmail(smtpClient, recipients))

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(fmt.Errorf("something went wrong while initializating the serve %w", err))
	}
}

func handleSendEmail(smtpClient smtpclient.Client, recipients []string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var email Email

		err := json.NewDecoder(r.Body).Decode(&email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if email.Body == "" || email.Subject == "" {
			http.Error(w, "Empty members", http.StatusBadRequest)
			return
		}

		message := smtpclient.BuildMessage(recipients, email.Subject, email.Body)

		err = smtpClient.SendEmail(recipients, message)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}

}
