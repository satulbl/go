package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30
	// ch <- 40 
	// Когда буфер полон компилятор останавливает выполнения кода и ждет 
	// пока другая горутина прочитает канал и освободит буфер, 
	// такого не происходит и программа блокируется навсегда - deadlock

	fmt.Printf("Ёмкость: %d\n", cap(ch))
	fmt.Printf("Длина: %d\n", len(ch))
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
