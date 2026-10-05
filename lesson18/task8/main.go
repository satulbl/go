package main

import "fmt"

func main() {
	var isEvenChClosed, isOddChClosed = false, false
	evenCh := make(chan int, 5)
	oddCh := make(chan int, 5)

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				evenCh <- i
			} else {
				oddCh <- i
			}
		}

		close(evenCh)
		close(oddCh)
	}()

	for {
		if isEvenChClosed && isOddChClosed {
			break
		}

		select {
		case msg, isClosed := <-evenCh:
			if !isClosed {
				isEvenChClosed = true
			} else {
				fmt.Printf("Even: %d\n", msg)
			}

		case msg, isClosed := <-oddCh:
			if !isClosed {
				isOddChClosed = true
			} else {
				fmt.Printf("Odd: %d\n", msg)
			}
		}
	}
}
