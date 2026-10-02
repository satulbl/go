package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Printf("Ёмкость: %d\n", cap(ch))
	fmt.Printf("Длина: %d\n", len(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Printf("Ёмкость: %d\n", cap(ch))
	fmt.Printf("Длина: %d\n", len(ch))
	fmt.Println(<-ch)
	fmt.Printf("Ёмкость: %d\n", cap(ch))
	fmt.Printf("Длина: %d\n", len(ch))
}
