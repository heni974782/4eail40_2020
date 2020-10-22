package main

import (
    "fmt"
	"math"
)

// Implement types Rectangle, Circle and Shape
type Rectangle struct{
	width, length float64
}

type Circle struct{
	radius float64
}

type Shape interface {
	Surface() float64
}

func (r Rectangle) Surface() float64 {
    return r.width * r.length
}

func (c Circle) Surface() float64 {
    return math.Pi * c.radius * c.radius
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Square of width %v and length %v", r.width, r.length)
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle of radius %v",c.radius)
}



func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}
	
	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		 
	}
}
}
