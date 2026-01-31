package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipeint := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formatedMsg := fmt.Sprintf("To: %s\r\nSubject: Hi %s\r\n\r\n%s\r\n", recipeint.Email, recipeint.Name, "just testing our email service")
		// msg := []byte(formatedMsg)

		msg, err := executeTemplate(recipeint)
		if err != nil {
			fmt.Printf("Worker: %d, Error parsing template for %s\n", id, recipeint.Email)
			//todo: Proper handling needed- add to dlq
			continue
		}

		fmt.Printf("worker %d, sending email to recipeint %s\n", id, recipeint.Name)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "saksham@example.com", []string{recipeint.Email}, []byte(msg))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("worker %d, sent email to recipeint %s\n", id, recipeint.Name)
		time.Sleep(50 * time.Millisecond)
	}
}
