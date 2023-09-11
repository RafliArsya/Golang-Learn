package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	runtime.GOMAXPROCS(4)
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)
		wg.Add(1)
		fmt.Println(time.Since(start))
		go doPrint(&wg, data)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
	var input string
	fmt.Scanln(&input)
}
