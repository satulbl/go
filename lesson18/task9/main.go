package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	messages := make(chan string)

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		stopTimer := time.After(2 * time.Second)

		for {
			select {
			case <-ticker.C:
				messages <- "ping"
			case <-stopTimer:
				close(done)
				return
			}
		}
	}()

	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		case msg := <-done:
			fmt.Println("Конец")
			fmt.Println(msg)
			return
		}
	}
}
