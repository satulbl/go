package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)

	val1, ok1 := <-ch
	fmt.Println(val1, ok1)

	val2, ok2 := <-ch
	fmt.Println(val2, ok2)

	val3, ok3 := <-ch
	fmt.Println(val3, ok3)

	val4, ok4 := <-ch
	fmt.Println(val4, ok4)
	// будем получать false и zero value (0) так как канал закрыт чтение безопасное, если не закрыть получим deadlock
}
