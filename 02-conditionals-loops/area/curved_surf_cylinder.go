package main

import (
	"fmt"
	"math"
)



type CurvedAreaCylinder struct {
	radius, height float64
}

func (c *CurvedAreaCylinder) CurvedArea() {
	curvedArea := 2 * math.Pi * c.radius * c.height
	fmt.Printf("The curved surface area of a cylinder is: %.2f\n", curvedArea)
}
