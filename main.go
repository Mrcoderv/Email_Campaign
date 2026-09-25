package main

import (
	"fmt"
	"time"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	fmt.Println("welcome to email sender")
	recipientChannel := make(chan Recipient)
	// dead lock situation 1/// here he worker are ready but producer is not ready
	// emailworker(1,recipientChannel)
	// loadRecipients("./Emails.csv", recipientChannel)
	// here the producer are ready but the worker are not ready ,
	// loadRecipients("./Emails.csv", recipientChannel)
	// emailworker(1,recipientChannel)
	// so we need to run both at the same time .using the go routine , concurently
	go func() {
		loadRecipients("./Emails.csv", recipientChannel)
	}()
	// or
	workerCount := 5
	for i := 1; i <= workerCount; i++ {
		go emailworker(i, recipientChannel)
	}

	go emailworker(1, recipientChannel)
	time.Sleep(3 * time.Second) // to make the wait the main proggram for the finishing the go routine task
	//   this is not the best method 

	fmt.Println("welcome to email sender")
}
