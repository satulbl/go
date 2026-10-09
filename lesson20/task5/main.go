package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	Title string `json:"title"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

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

	var p []Post
	errUnm := json.Unmarshal(body, &p)
	if errUnm != nil {
		fmt.Println(errUnm)
		return
	}

	fmt.Printf("Длина: %d\nПервый: %s\nВторой: %s\n", len(p), p[0].Title, p[len(p)-1].Title)
}
