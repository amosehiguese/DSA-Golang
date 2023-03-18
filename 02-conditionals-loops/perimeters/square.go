package main

import (
	"fmt"
)



type Square struct {
	sides float64
}

func (s *Square) Perimeter() {
	perimeter := 4 * s.sides
	fmt.Printf("Area of the Square is: %.2f\n", perimeter)
}

