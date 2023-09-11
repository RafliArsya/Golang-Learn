package main

import (
	"fmt"
)

func dorelational(val1 int, val2 int) {
	var result = (val1 == val2)
	fmt.Printf("Apakah dua nilai sama = %t\n", result)
}
