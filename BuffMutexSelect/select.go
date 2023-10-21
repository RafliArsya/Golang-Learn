package main

import (
	"fmt"
	"runtime"
)

func getAvg(nr []int, ch chan float64) {
	var sum = 0
	for _, v := range nr {
		sum += v
	}
	ch <- float64(sum) / float64(len(nr))
}

func getMax(nr []int, ch chan int) {
	var max = nr[len(nr)-1]
	for _, v := range nr {
		if max < v {
			max = v
		}
	}
	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)
	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)
	var ch1 = make(chan float64)
	go getAvg(numbers, ch1)
	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg\t: %.2f\n", avg)
		case max := <-ch2:
			fmt.Printf("Max\t: %d\n", max)
		}
	}
}
