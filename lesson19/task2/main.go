package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	u := User{Name: "Yusuf", Age: 20, Email: "y.norov@gmail.com"}

	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
