package main

import (
	"encoding/csv"
	"os"
)

func loadRecipients(filePath string, ch chan Recipient) error {
	file, e := os.Open(filePath)
	if e != nil {
		return e

	}
	defer file.Close()
	r := csv.NewReader(file)

	records, err := r.ReadAll()

	if err != nil {
		return err

	}

	for _, record := range records[1:] { //  do not need to assume the number of column  index  0is the head title
		//fmt.Println(record)
		// send --> consumer ---. channel
		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}
	return nil
}
