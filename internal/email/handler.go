package email

import (
	"net/http"

	httperror "github.com/AlvaroAriel/HTTP-SMTPClient/internal/error"
	"github.com/AlvaroAriel/HTTP-SMTPClient/internal/server"
	smtpclient "github.com/AlvaroAriel/HTTP-SMTPClient/smptclient"
)

func HandleSendEmail(smtpClient smtpclient.Client) http.HandlerFunc {
	recipients := getRecipient()

	return func(w http.ResponseWriter, r *http.Request) {

		email, err := server.DecodeJSON[Email](r)

		if err != nil {
			httperror.JSONError(w, httperror.InvalidJSON())
			return
		}

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
