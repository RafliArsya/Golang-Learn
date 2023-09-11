package main

import (
	"errors"
	"fmt"
	"strings"
)

func catched() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Input cannot be empty")
	}
	return true, nil
}

func main() {
	defer catched()
	var name string
	fmt.Print("Type your name\t: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Halo", name)
	} else {
		defer fmt.Println("Executed")
		panic(err.Error())
	}

}
