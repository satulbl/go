package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	person := Person{
		Name: "Rakhim",
		Age:  19,
		Address: Address{
			City:   "Dushanbe",
			Street: "Rudaki",
		}}

	data, err := json.MarshalIndent(person, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", string(data))

	var p Person

	errParse := json.Unmarshal([]byte(data), &p)
	if errParse != nil {
		fmt.Println(errParse)
		return
	}

	fmt.Printf("%+v\n", p)
}
