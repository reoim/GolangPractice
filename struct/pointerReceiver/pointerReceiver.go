package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(rect.width, area)
	area2 := rect.area2()
	println(rect.width, area2)

}
