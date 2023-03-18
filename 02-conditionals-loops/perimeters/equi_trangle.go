package main

import (
	"fmt"
)



type EquiTriangle struct {
	sides float64
}

func (e *EquiTriangle) Perimeter() {
	perimeter := 3 * e.sides
	fmt.Printf("Perimeter of the EquiTriangle is: %.2f\n", perimeter)
}

