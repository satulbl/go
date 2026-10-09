package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Post struct {
	Title string `json:"title"`
}

func main() {
	userId := 3
	baseUrl := "https://jsonplaceholder.typicode.com/posts"

	base, err := url.Parse(baseUrl)
	if err != nil {
		fmt.Println(err)
	}

	params := url.Values{}
	params.Add("userId", fmt.Sprint(userId))

	base.RawQuery = params.Encode()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, errR := client.Get(base.String())
	if errR != nil {
		fmt.Println(errR)
		return
	}

	defer resp.Body.Close()

	body, errB := io.ReadAll(resp.Body)
	if errB != nil {
		fmt.Println(errB)
		return
	}

	var p []Post
	errUnm := json.Unmarshal(body, &p)
	if errUnm != nil {
		fmt.Println(errUnm)
		return
	}

	fmt.Printf("Длина: %d\n", len(p))

	for postIndex, post := range p {
		fmt.Printf("%d: %s\n", postIndex+1, post.Title)
	}
}
