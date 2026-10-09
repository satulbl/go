package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	Title string `json:"title"`
}

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func fetchHeader(url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Student", "Farrukh")

	resp, errReq := client.Do(req)
	if errReq != nil {
		return errReq
	}

	defer resp.Body.Close()

	body, errB := io.ReadAll(resp.Body)
	if errB != nil {
		return errB
	}

	fmt.Println(string(body))
	return nil
}

func main() {
	url := "https://httpbin.org/headers"

	err := fetchHeader(url)
	if err != nil {
		fmt.Println(err)
		return
	}
}
