package main

import (
	"fmt"
)

func tryRead(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Printf("Прочитано сообщение: %q\n", msg)
	default:
		fmt.Println("Пусто")
	}
}

func main() {
	ch := make(chan string, 1)

	ch <- "Привет, Go!"

	tryRead(ch)

	tryRead(ch)

	tryRead(ch)
}
