package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"` //не обязательно, не привиредлив к регистру "AUThor" - тоже сработает, но лучше явно писать
	Year   float64 `json:"year"`
}

func main() {
	var b Book
	jsonStr := `{"title":"Go для начинающих","author":"Иван Петров","year":2023}`

	err := json.Unmarshal([]byte(jsonStr), &b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", b)
}
