package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circumference() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumference() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumference())
}

func main() {
	shapes := []shape{
		circle{radius: 2.5},
		square{length: 4.2},
		circle{radius: 7.8},
		square{length: 19.9},
	}

	for _, v := range shapes {
		fmt.Println("********")
		printShapeInfo(v)
		fmt.Println("********")
	}
}
