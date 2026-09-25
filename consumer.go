package main

import "fmt"

func emailworker(is int, ch chan Recipient) {
	for recipient := range ch {
		fmt.Print(is, recipient, "\n")
	}

}
