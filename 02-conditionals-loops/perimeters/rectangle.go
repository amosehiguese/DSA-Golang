package main

import (
	"fmt"
)



type Rectangle struct {
	length, width float64
}

func (r *Rectangle) Perimeter() {
	perimeter := 2 * (r.length + r.width)
	fmt.Printf("Perimeter of the Rectangle is: %.2f\n", perimeter)
}
