package main

import (
	"fmt"
	box "golang_united_school_homework"
)

func main () {
	circle1 := &box.Circle{Radius: 20}
	circle2 := &box.Circle{Radius: 30}
	circle3 := &box.Circle{Radius: 5}
	circle4 := &box.Circle{Radius: 2}
	triangle := &box.Triangle{Side: 30}
	rectangle := &box.Rectangle{Height: 10, Weight: 20}
	b := box.NewBox(6)
	_ = b.AddShape(circle1)
	_ = b.AddShape(circle2)
	_ = b.AddShape(circle3)
	_ = b.AddShape(circle4)
	_ = b.AddShape(triangle)
	_ = b.AddShape(rectangle)

	 b.RemoveAllCircles()

	 fmt.Println(b.GetLen())
	 
}