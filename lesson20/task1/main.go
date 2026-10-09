package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	fmt.Printf("%v\n", resp.StatusCode)
	fmt.Printf("%v\n", string(body))
	fmt.Printf("%v\n", resp.Header.Get("Content-Type"))
}
