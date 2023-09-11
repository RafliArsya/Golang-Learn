package main

import "fmt"

func domaps() {
	var chicken = map[string]int{} //chicken := map[string]int{}
	chicken["january"] = 50
	chicken["february"] = 40
	fmt.Println("Data januari adalah\t", chicken["january"])
	fmt.Println("Semua data\t", chicken)
}
