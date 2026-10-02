package main

import "fmt"

func sendUsers(users []string, ch chan string) {
	for _, user := range users {
		ch <- user
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	users := []string{"Аня", "Борис", "Вика", "Данил", "Ева"}
	usersCopy := make([]string, 0, len(users))

	go sendUsers(users, ch)

	for user := range ch {
		usersCopy = append(usersCopy, user)
	}

	fmt.Println(len(usersCopy), usersCopy)
}
