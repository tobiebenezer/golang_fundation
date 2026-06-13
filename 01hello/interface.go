package main

import (
	"fmt"
	"math"
)

type Shape interface {

	getPerimeter() uint

	getArea() uint
}

type Rectangle struct {
	l uint
	b uint
}

func (this Rectangle) getPerimeter()  uint{
	return (this.l + this.b) * 2
}

func (this Rectangle) getArea() uint {
	return  this.b * this.l
}


type Square struct {
	width uint
}

func (this Square) getPerimeter() uint {
	return this.width * 4
}

func (this Square) getArea() uint{
	return uint(math.Pow(float64(this.width), 2))
}

func main3() {
	var shapes []Shape = []Shape{Rectangle{5, 6}, Square{5}}

	total_permeter := uint(0)
	for _, shape := range shapes{
		total_permeter += shape.getPerimeter()
	}

	fmt.Println(total_permeter)
}