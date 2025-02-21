package main

import (
	"fmt"
	"log"

	smtpclient "github.com/AlvaroAriel/HTTP-SMTPClient/smptclient"
)

func main() {

	envPath := ".env"
	recipients := []string{"insert some email"}
	subject := "sub"
	body := "some text"
	message := smtpclient.BuildMessage(recipients, subject, body)

	client, err := smtpclient.BuildClient(envPath)

	if err != nil {
		log.Fatal(err)
	}

	err = client.SendEmail(recipients, message)

	if err != nil {
		fmt.Println("Failed sending email: ", err)
	}

}
