package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Воркер <%d> закончил\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go worker(i, &wg)
	}

	wg.Wait()
}
