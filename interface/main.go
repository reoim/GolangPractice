package main

import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// interface instance of Rect type
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// interface instance of Circle type
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}

func main() {
	r := Rect{10, 20}
	c := Circle{10}

	showArea(r, c)

	var a interface{} = 1

	i := a       // a와 i 는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1

	println(i) // 포인터주소 출력
	println(j)
}
