package structsandinterface

import (
	// "fmt"
	"math"
)




type Shape interface {
	Area() float64
}


type Rectangle struct {
	length float64
	breadth float64
}

func(rectangle Rectangle) Perimeter() float64 {
	return 2*(rectangle.length+rectangle.breadth)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.length*rectangle.breadth
}

type Circle struct {
	radius float64
}

func (r Circle)Area()float64 {
	return math.Pi*r.radius*r.radius
}
func (r Circle)Perimeter() float64 {
	return 2*math.Pi*r.radius
}