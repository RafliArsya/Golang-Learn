package main

import "fmt"

func doarray() {
	var fruits = [4]string{"apple", "grapes", "banana", "melons"}
	fmt.Println("Panjang data\t", len(fruits))
	fmt.Println("Isi data\t", fruits)
	fmt.Println("Isi data pada index 1\t", fruits[1])
}
