package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	InStock bool    `json:"inStock"`
}

func main() {
	pr := Product{ID: 214214, Name: "Parrot", Price: 400, InStock: true}

	data, err := json.Marshal(pr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
