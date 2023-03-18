package main

import (
	"fmt"
)



type Rhombus struct {
	sides float64
}

func (s *Rhombus) Perimeter() {
	perimeter := 4 * s.sides
	fmt.Printf("Area of the Rhombus is: %.2f\n", perimeter)
}

