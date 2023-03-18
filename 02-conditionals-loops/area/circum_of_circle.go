package main

import (
	"fmt"
	"math"
)



type CircumCircle struct {
	radius float64
}

func (c *CircumCircle) Circumference() {
	circumference := 2 * math.Pi * c.radius
	fmt.Printf("Area of the CircumCircle is: %.2f\n", circumference)
}
