package golang_united_school_homework

import (
	"errors"
	"fmt"
	"reflect"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("box capacity more then limit")
	}
	b.shapes = append(b.shapes, shape)
	return nil
	// panic("implement me")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes) - 1 { return nil, errors.New("index went out of the range")}
	if b.shapes[i] == nil { return nil, errors.New("shape by index doesn't exist") }
	return b.shapes[i], nil
	// panic("implement me")

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var t []Shape

	s, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	
	t = append(t, b.shapes[0:i]...)
	t = append(t, b.shapes[i+1:len(b.shapes)]...)
	b.shapes = t
	return s, err
	// panic("implement me")
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var t []Shape

	s, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	
	t = append(t, b.shapes[0:i]...)
	t = append(t, shape)
	t = append(t, b.shapes[i+1:len(b.shapes)]...)
	b.shapes = t
	return s, err
	// panic("implement me")

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}

	return sum
	// panic("implement me")

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}

	return sum
	// panic("implement me")

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var i, j int

	for {
		if i >= len(b.shapes) { break }
		if reflect.TypeOf(b.shapes[i]).Name() != "Circle" { 
			i++
		} else {
			_, err := b.ExtractByIndex(i)
			if err != nil {
				fmt.Printf("Error Extraction by index in RemoveAllCircles method: %s", err.Error())
			}
			j++
		}
	}

	if j > 0 {
		return nil
	} 

	return errors.New ("circles are not exist in the list")
	// panic("implement me")

}
