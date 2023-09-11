package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name": "John Wick", "Age": 27}`
	var jsonData = []byte(jsonString)
	var data map[string]interface{}
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("User\t:", data["Name"])
	fmt.Println("Age\t:", data["Age"])
}
