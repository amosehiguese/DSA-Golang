package main

import (
	"fmt"
	"math"
)



type Cone struct {
	height, radius float64
}

func (c *Cone) Volume() {
	volume := 0.33 * math.Pi * math.Pow(c.radius,2) * c.height
	fmt.Printf("Volume of the Cone is: %.2f\n", volume)
}

