package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Company struct {
	Name string `json:"name"`
}

type User struct {
	Name    string  `json:"name"`
	Company Company `json:"company"`
}

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func GetUsers(url string) ([]User, error) {
	resp, errResp := client.Get(url)
	if errResp != nil {
		return nil, errResp
	}

	defer resp.Body.Close()

	var users []User
	err := json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Status)
	return users, nil
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users"

	users, err := GetUsers(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(users))
	for _, user := range users {
		fmt.Printf("Имя: %s, Название компании: %s\n", user.Name, user.Company.Name)
	}
}
