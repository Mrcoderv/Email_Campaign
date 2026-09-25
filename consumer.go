package main

import (
	"fmt"
	"net/smtp"
	"sync"
)

func emailworker(i int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done() // finished task

	for recipient := range ch {
		smtpSender := "raghavpanthi@yahoo.com"
		smtpHost := "localhost"
		smtpPort := 1025

		// formattedMsg := fmt.Sprintf(
		// 	"To: %s\r\nSubject: test email\r\n\r\njust testing", recipient.Email)

		msg, er := exeTemplate(recipient) //  in the main.go file
		if er != nil {
			fmt.Println("worker: errror %s", recipient.Email)
			continue // continue next email
		}
		err := smtp.SendMail(
			fmt.Sprintf("%s:%d", smtpHost, smtpPort),
			nil,
			smtpSender,
			[]string{recipient.Email},
			[]byte(msg),
		)

		if err != nil {
			fmt.Println("Email sending error:", err)
			continue
		}
		fmt.Printf("channel %d send  Email sent to: %s\n", i, recipient.Email)

	}

}
