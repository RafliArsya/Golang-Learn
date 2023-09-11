package main

import "fmt"

func main() {
	var firstname string = "Leo"
	const address = "Jakarta"
	var beratpecahan float64 = 0.25
	var berat = 50

	var A int = 4
	var pA *int = &A

	fmt.Printf("Halo %s di %s!\n", firstname, address)
	fmt.Printf("Berat kamu adalah %g dari %d kg dan %g kg\n", beratpecahan+float64(berat), berat, beratpecahan)
	fmt.Printf("A = %d\nAddress A = %p.\npA = %p\nPointer pA = %d\nAddress pA = %p\n", A, &A, pA, *pA, &pA)
}
