package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Halo")
	os.Exit(1)
	fmt.Println("Selamat datang")
}
