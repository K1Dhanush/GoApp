package main

import "fmt"

//interface
type Rectangle interface {
	Area() float64 //methods and return type
	Perimeter() float64
}

//type struct just like a class
type values struct {
	length  int
	breadth int
}

func (l values) Area() float64 {
	return float64(l.length * l.breadth)
}

func (l values) Perimeter() float64 {
	return float64(2 * (l.length * l.breadth))
}
func main() {
	var t Rectangle
	t = values{4, 5}
	fmt.Println("Area of a rectangle ", t.Area())
	fmt.Println("Perimeter of a rectangle ", t.Perimeter())
}
