package main

import (
	"fmt"
)



type Isosceles struct {
	base,height float64
}

func (t *Triangle) Area() {
	area := 0.5 * t.base * t.height
	fmt.Printf("Area of the Isosceles is: %.2f\n", area)
}

