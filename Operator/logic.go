package main

import (
	"fmt"
)

func dological(val1 bool, val2 bool) {
	var result = (val1 && val2)
	fmt.Printf("Apakah dua nilai sama = %t\n", result)
}
