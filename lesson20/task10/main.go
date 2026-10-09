package main

import (
	"bytes"
	"encoding/json"
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
	Title  string `json:"title"`
	Body   string `json:"body"`
	Id int `json:"id"`
}

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func CreatePost(p NewPost, url string) (*CreatedPost, error) {
	payload, errConvert := json.Marshal(p)
	if errConvert != nil {
		return nil, errConvert
	}

	resp, errResp := client.Post(url, "application/json", bytes.NewBuffer(payload))
	if errResp != nil {
		return nil, errResp
	}

	defer resp.Body.Close()

	var post CreatedPost
	errDecode := json.NewDecoder(resp.Body).Decode(&post)
	if errDecode != nil {
		return nil, errDecode
	}

	return &post, nil
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	post := NewPost{Title: "Пост", Body: "Не знаю что тут должно быть", UserId: 454}

	resp, err := CreatePost(post, url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Id)
	fmt.Println(resp.Body)
	fmt.Println(resp.Title)
}
