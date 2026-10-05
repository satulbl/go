package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	for i := 1; i <= 5; i++ {
		select {
		case ch <- i:
			fmt.Printf("Отправили %d\n", i)
		default:
			fmt.Printf("Пропускаем <%d>\n", i)
		}
	}
}
