package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(till int, message string, start time.Time) {
	for i := 0; i < till; i++ {
		elapsed := time.Since(start)
		fmt.Println(elapsed, "\t:", (i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	start := time.Now()
	print(5, "Hai Apa", start)
	print(5, "Kabar Mantan", start)
	start = time.Now()
	go print(5, "go Hai Apa", start)
	go print(5, "go Kabar Mantan", start)

	var input string
	fmt.Scanln(&input)
}
