package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Type some number\t: ")
	fmt.Scanln(&input)
	var number uint64
	var err error
	number, err = strconv.ParseUint(input, 10, 64)
	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "isn't number")
		fmt.Println(err.Error())
	}
}
