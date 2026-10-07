package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Login    string
	Password string `json:"-"`
	Balance  float64
}

func main() {
	ac := Account{Login: "r.yunusov@gmail.com", Password: "test3627", Balance: 545.5}

	data, err := json.Marshal(ac)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
