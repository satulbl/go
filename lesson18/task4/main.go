package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("Данных пока нет")
	}

}
