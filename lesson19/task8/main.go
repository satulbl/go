package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"` // и с json и без count - zero value
}

func main() {
	var item Item
	jsonStr := `{"name":"Ноутбук","price":4500.5}`

	err := json.Unmarshal([]byte(jsonStr), &item)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", item)
}
