package main

import (
	"fmt"
	"math"
)



type Cylinder struct {
	height, radius float64
}

func (c *Cylinder) Volume() {
	volume := math.Pi * math.Pow(c.radius,2) * c.height
	fmt.Printf("Volume of the Cylinder is: %.2f\n", volume)
}

