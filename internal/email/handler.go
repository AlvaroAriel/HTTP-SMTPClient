package email

import (
	"encoding/json"
	"net/http"

	httperror "github.com/AlvaroAriel/HTTP-SMTPClient/internal/error"
	smtpclient "github.com/AlvaroAriel/HTTP-SMTPClient/smptclient"
)

func HandleSendEmail(smtpClient smtpclient.Client) http.HandlerFunc {
	recipients := getRecipient()

	return func(w http.ResponseWriter, r *http.Request) {

		var email Email
		err := json.NewDecoder(r.Body).Decode(&email)

		if err != nil {
			httperror.JSONError(w, httperror.InvalidJSON())
			return
		}

		defer r.Body.Close()

		if email.Body == "" || email.Subject == "" {
			httperror.JSONError(w, httperror.EmptyField())
			return
		}

		message := smtpclient.BuildMessage(recipients, email.Subject, email.Body)

		err = smtpClient.SendEmail(recipients, message)

		if err != nil {
			httperror.JSONError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}

}
