package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users/1"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}

	defer resp.Body.Close()

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		fmt.Println("Ошибка чтения:", errRead)
		return
	}

	var u User
	errUnm := json.Unmarshal(body, &u)
	if errUnm != nil {
		fmt.Println(errUnm)
		return
	}

	fmt.Println(u.Name)
	fmt.Println(u.Email)
	fmt.Println(u.Phone)
}
