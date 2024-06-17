package main

import "fmt"

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{base: 10, height: 20}
	s := square{sideLength: 10}
	printArea(t)
	printArea(s)
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}