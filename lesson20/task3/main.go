package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Task struct {
	Title string `json:"title"`
}

func printTitle(resp *http.Response) {
	fmt.Println(resp.StatusCode)

	switch resp.StatusCode { // useless (*resp).StatusCode
	case http.StatusOK:
		body, errRead := io.ReadAll(resp.Body)
		if errRead != nil {
			fmt.Println("Ошибка чтения:", errRead)
			return
		}

		var t Task
		errUnm := json.Unmarshal(body, &t)
		if errUnm != nil {
			fmt.Println("Ошибка парсинга JSON:", errUnm)
			return
		}

		fmt.Println(t.Title)

	case http.StatusNotFound:
		fmt.Println("Задача не найдена")

	default:
		fmt.Printf("Непредвиденный статус-код: %d\n", resp.StatusCode)
	}
}

func main() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/5",
		"https://jsonplaceholder.typicode.com/todos/99999",
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, url := range urls {
		resp, errReq := client.Get(url)
		if errReq != nil {
			fmt.Println("Ошибка запроса", errReq)
			continue
		}

		printTitle(resp)
		resp.Body.Close()
	}
}
