//Area of a circle = pie * square(radius)

package main

import (
	"fmt"
	"math"
)



type Circle struct {
	radius float64
}

func (c *Circle) Area() {
	area := math.Pi * math.Pow(c.radius, 2)
	fmt.Printf("Area of the Circle is:%.2f\n", area)
}



