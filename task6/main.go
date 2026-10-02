package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	ch1 := make(chan int, 5)

	ch1 <- 10
	ch1 <- 20
	ch1 <- 30
	ch1 <- 40
	ch1 <- 50

	fmt.Println(time.Since(time1))

	ch2 := make(chan int, 5)
	time2 := time.Now()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(200 * time.Millisecond)
			<-ch2
		}
	}()

	for i := 0; i < 6; i++ {
		ch2 <- i
	}

	fmt.Println(time.Since(time2))
}
