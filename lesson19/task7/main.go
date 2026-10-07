package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p Person
	jsonStr := `{"name":"Алия","age":20,"city":"Душанбе"}`

	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", p)
}
