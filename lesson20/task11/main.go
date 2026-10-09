package main

import (
	"fmt"
	"net/http"
	"time"
)

type NewPost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

type CreatedPost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Id    int    `json:"id"`
}

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func DeletePost(url string) error {
	req, errReq := http.NewRequest(http.MethodDelete, url, nil)
	if errReq != nil {
		return errReq
	}

	resp, errR := client.Do(req)
	if errR != nil {
		return errR
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)
	return nil
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	err := DeletePost(url)
	if err != nil {
		fmt.Println(err)
	}
}
