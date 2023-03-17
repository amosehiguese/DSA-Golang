//Area of a circle = pie * square(radius)

package main

import (
	"fmt"
	"math"
)



type Circle struct {
	Radius float64
}

func (c *Circle) Area() {
	Area := math.Pi * math.Pow(c.Radius, 2)
	fmt.Printf("Area of the Circle is:%.2f\n", Area)
}

func main() {
	c := Circle{9}
	c.Area()
}

