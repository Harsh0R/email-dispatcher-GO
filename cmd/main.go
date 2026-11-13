package main

import (
	"fmt"
	"sync"
)

type Recipient struct {
	Name    string
	Email   string
	Content string
}

func main() {

	recepientChannel := make(chan Recipient)

	var wg sync.WaitGroup

	go func() {
		if err := loadRecipient("./dummy_users.csv", recepientChannel); err != nil {
			fmt.Println("err -> ", err)
		}
	}()

	workerCount := 5
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recepientChannel, &wg)
	}
	wg.Wait()
}
