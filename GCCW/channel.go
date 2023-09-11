package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	var message = make(chan string)
	var sendTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		message <- data
	}
	start := time.Now()
	go sendTo("john")
	go sendTo("ethan")
	go sendTo("ren")
	var message1 = <-message
	elapsed := time.Since(start)
	fmt.Println(elapsed, ":", message1)
	elapsed = time.Since(start)
	var message2 = <-message
	fmt.Println(elapsed, ":", message2)
	elapsed = time.Since(start)
	var message3 = <-message
	fmt.Println(elapsed, ":", message3)

	var input string
	fmt.Scanln(&input)
}
