package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
	Price  float64
}

func main() {
	b := Book{Title: "Book", Author: "Blok", Year: 1960, Price: 150}

	data, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
