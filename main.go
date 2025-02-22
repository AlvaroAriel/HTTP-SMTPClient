package main

import (
	"fmt"
	"log"

	smtpclient "github.com/AlvaroAriel/HTTP-SMTPClient/smptclient"
)

func main() {

	recipients := []string{"some mails"}
	subject := "sub"
	body := "some text"
	message := smtpclient.BuildMessage(recipients, subject, body)

	client, err := smtpclient.BuildClient()

	if err != nil {
		log.Fatal(err)
	}

	err = client.SendEmail(recipients, message)

	if err != nil {
		fmt.Println("Failed sending email: ", err)
	}

}
