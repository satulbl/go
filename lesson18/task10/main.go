package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int
	var wg sync.WaitGroup
	results := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- i
		}()
	}

	wg.Wait()
	go func() {
		close(results)
	}()
	
	for v := range results {
		fmt.Println(v)
		sum += v
	}

	fmt.Println(sum)
}
