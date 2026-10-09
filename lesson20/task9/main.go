package main

import (
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	Title string `json:"title"`
}

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func fetchStatus(url string) error {
	resp, errR := client.Get(url)
	if errR != nil {
		return errR
	}

	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Println("<200>: успех")
	case http.StatusCreated:
		fmt.Println("<201>: успех")
	case http.StatusNotFound:
		fmt.Println("<404>: ошибка клиента")
	case http.StatusInternalServerError, http.StatusServiceUnavailable:
		fmt.Printf("<%d>: ошибка сервера\n", resp.StatusCode)
	}

	return nil
}

func main() {
	url := "https://httpbin.org/status"
	statuses := []int{200, 201, 404, 500, 503}

	for _, status := range statuses {
		baseUrl := fmt.Sprintf("%s/%d", url, status)
		err := fetchStatus(baseUrl)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
