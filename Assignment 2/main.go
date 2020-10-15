package main

import "fmt"

type shape interface {
	getArea() float64
}
type square struct {
	side float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {
	square1 := square{}
	square1.side = 4
	printArea(square1)
	triangle1 := triangle{}
	triangle1.base = 4
	triangle1.height = 4
	printArea(triangle1)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())

}
func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {
	return (0.5) * t.base * t.height
}
