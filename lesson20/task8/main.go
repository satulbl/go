package main

import (
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	Title string `json:"title"`
}

var (
	clientGood = &http.Client{
		Timeout: 5 * time.Second,
	}
	clientBad = &http.Client{
		Timeout: 1 * time.Second,
	}
)

func fetchDelay(url string, client *http.Client) error {
	req, err := client.Get(url)
	if err != nil {
		return err
	}

	fmt.Println(req.Status)
	return nil
}

func main() {
	url := "https://httpbin.org/delay/2"

	err1 := fetchDelay(url, clientGood)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	err2 := fetchDelay(url, clientBad)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
