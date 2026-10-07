package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string `json:"name"`
	Grade   int    `json:"grade"`
	Comment string `json:"comment,omitempty"`
}

func main() {
	st1 := Student{Name: "Farrukh", Grade: 5}
	st2 := Student{Name: "Farrukh", Grade: 5, Comment: "comment thst"}

	data1, err1 := json.Marshal(st1)
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(string(data1))

	data2, err2 := json.Marshal(st2)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(string(data2))
}
