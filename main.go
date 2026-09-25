package main

import (
	"fmt"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	fmt.Println("welcome to email sender")
	recipientChannel := make(chan Recipient)

	loadRecipients("./Emails.csv", recipientChannel)
}
