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
	for recepient := range ch {

		smtpHost := "localhost"
		smtpPort := "1025"
		// formatMsg := fmt.Sprintf("To: %s\r\nSubject: Hii %s, This is Test Mail.\r\n\r\n%s\r\n", recepient.Email, recepient.Name, recepient.Content)
		// msg := []byte(formatMsg)

		msg, err := executeTemplate(recepient)
		if err != nil {
			fmt.Printf("Worker id: %d, Error parsing for %s", id, recepient.Email)
			continue
		}

		fmt.Printf("worker %d sending email to %s", id, recepient.Email)
		// err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "harshradadiya9999@gmail.com", []string{recepient.Email}, msg)
		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "harshradadiya9999@gmail.com", []string{recepient.Email}, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("worker %d sent email to %s", id, recepient.Email)

	}
}
