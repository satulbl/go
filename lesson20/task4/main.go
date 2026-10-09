package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Task struct {
	Completed bool `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, errReq := client.Get(url)
	if errReq != nil {
		fmt.Println(errReq)
		return
	}

	defer resp.Body.Close()

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		fmt.Println(errRead)
		return
	}

	var todos []Task
	errUnm := json.Unmarshal(body, &todos)
	if errUnm != nil {
		fmt.Println(errUnm)
		return
	}

	var completedCount, notCompletedCount int
	for _, t := range todos {
		if t.Completed {
			completedCount++
		} else {
			notCompletedCount++
		}
	}

	fmt.Println(len(todos))
	fmt.Printf("Закончено: %d\n", completedCount)
	fmt.Printf("Не закончено: %d\n", notCompletedCount)
}
