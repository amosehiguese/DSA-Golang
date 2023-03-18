package main

import (
	"fmt"
)



type Triangle struct {
	height float64
	base float64
}

func (t *Triangle) Area() {
	area := 0.5 * t.base * t.height
	fmt.Printf("Area of the Triangle is: %.2f\n", area)
}

