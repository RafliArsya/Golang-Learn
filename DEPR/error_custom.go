package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Input cannot be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("Type your name\t: ")
	fmt.Scanln(&name)

	/*valid, err := validate(name)
	if valid {
		fmt.Println("Halo", name)
	} else {
		fmt.Println(err.Error())
	}*/
	if valid, err := validate(name); valid {
		fmt.Println("Halo", name)
	} else {
		fmt.Println(err.Error())
	}

}
