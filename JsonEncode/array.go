package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `[
		{"Name": "John Wick", "Age": 27},
		{"Name": "Ethan Winter", "Age": 25}
	]`
	var data []User
	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 2; i++ {
		fmt.Println("User\t:", data[i].FullName)
		fmt.Println("Age\t:", data[i].Age)
	}
}
