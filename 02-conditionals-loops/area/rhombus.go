package main

import (
	"fmt"
)



type Rhombus struct {
	base,height float64
}

func (r *Rhombus) Area() {
	area := 0.5 * r.base *r.height
	fmt.Printf("Area of the Rhombus is: %.2f\n", area)
}

