package main

import (
	"fmt"
	"math"
)



type Equilateral struct {
	length float64
}

func (e *Equilateral) Area() {
	area := (math.Sqrt(3)/4) * math.Pow(e.length,2)
	fmt.Printf("Area of the Equilateral is: %.2f\n", area)
}

