package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up Authentication
	auth := smtp.PlainAuth("", "jayanta@zenius.one", "Panthera25@Tigris", "smtppro.zoho.in")

	// Connect to the server, Set the Sender and Recepient, and send the email
	to := []string{"bboyhotshot@gmail.com"}
	msg := []byte("To")
	err := smtp.SendMail("smtppro.zoho.in:465", auth, "hello@zenius.one", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
