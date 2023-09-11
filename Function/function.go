package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//basic function
	var names = []string{"Red", "Leo"}
	printMessage("Halo", names)

	//multiple return values
	var area, circum float64
	area, circum = calculate(3.5)
	var s_area = strconv.FormatFloat(area, 'f', 5, 64)
	var s_circum = strconv.FormatFloat(circum, 'f', 5, 64)
	printMessage("Area and Circumference of Circle are\t", []string{s_area, s_circum})

	//variadic argument
	var avg = calculate_avg(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	printMessage("Rata-rata:\t", []string{strconv.FormatFloat(avg, 'f', 2, 64)})

	//closure or local function
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var mynum = []int{4, 3, 2, 1, 3, 5}
	var min, max = getMinMax(mynum)
	printMessage("Min Max\t", []string{strconv.FormatInt(int64(min), 10), strconv.FormatInt(int64(max), 10)})

	//Callback
	var hasil = filter([]string{"ini", "data"}, func(each string) bool {
		return each == "ini"
	})
	fmt.Println(hasil)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference
}

func calculate_avg(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
