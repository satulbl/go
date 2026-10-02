package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)

	go square(3, ch)

	fmt.Println(<-ch)
}
