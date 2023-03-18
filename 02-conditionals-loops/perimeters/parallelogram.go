package main

import (
	"fmt"
)



type Parallelogram struct {
	a, b float64
}

func (p *Parallelogram) Perimeter() {
	perimeter := 2 * (p.a + p.b)
	fmt.Printf("Area of the Parallelogram is: %.2f\n", perimeter)
}

