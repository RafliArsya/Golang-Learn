package main

import (
	"fmt"
)

func doswitch(arga uint) {
	var point = arga
	switch point {
	case 10, 9:
		fmt.Println("Lulus dengan nilai sempurna")
	case 8, 7, 6, 5:
		fmt.Println("Lulus")
	case 4:
		fmt.Println("Hampir Lulus")
	default:
		fmt.Printf("Tidak lulus dengan nilai %d", point)
	}
}
