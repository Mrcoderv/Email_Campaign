package main

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"
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
	go func() { // producer
		loadRecipients("./Emails.csv", recipientChannel)
	}()
	// or
	var wg sync.WaitGroup
	workerCount := 5
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailworker(i, recipientChannel, &wg) //  calling consumer file functiom
	}
	//time.Sleep(3 * time.Second) // to make the wait the main proggram for the finishing the go routine task
	//   this is not the best method

	wg.Wait()
}
func exeTemplate(r Recipient) (string, error) {
	temp, err := template.ParseFiles("email.tmpl")

	if err != nil {
		return "", err

	}
	var tpl bytes.Buffer        //  buffer to store the output of the template.. execution
	err = temp.Execute(&tpl, r) // this execute the template and write the output to the buffer tpl
	if err != nil {
		return "", err

	}

	return tpl.String(), nil //

}
