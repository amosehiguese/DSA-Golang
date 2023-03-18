package main

import (
	"fmt"
)



type Isosceles struct {
	base,height float64
}

func (i *Isosceles) Area() {
	area := 0.5 * i.base *i.height
	fmt.Printf("Area of the Isosceles is: %.2f\n", area)
}

