package main

import (
	"fmt"
	"math"
)

type count interface {
	area() float64
	circ() float64
}

type rect struct {
	side float64
}

type circle struct {
	diameter float64
}

func (r rect) area() float64 {
	return math.Pow(r.side, 2)
}

func (r rect) circ() float64 {
	return r.side * 4
}

func (c circle) radius() float64 {
	return c.diameter * 0.5
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}

func (c circle) circ() float64 {
	return math.Pi * c.diameter
}

func main() {
	var shape count
	shape = rect{float64(10)}
	fmt.Println("luas persegi\t\t:", shape.area())
	fmt.Println("keliling persegi\t:", shape.circ())
	shape = circle{float64(14)}
	fmt.Println("jari-jari lingkaran\t:", shape.(circle).radius())
	fmt.Println("luas lingkaran\t\t:", shape.area())
	fmt.Println("keliling lingkaran\t:", shape.circ())
}
