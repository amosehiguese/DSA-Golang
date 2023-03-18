package main

import (
	"fmt"
	"math"
)



type Sphere struct {
	radius float64
}

func (s *Sphere) Volume() {
	volume := 1.33 * math.Pi * math.Pow(s.radius,3)
	fmt.Printf("Volume of the Sphere is: %.2f\n", volume)
}
