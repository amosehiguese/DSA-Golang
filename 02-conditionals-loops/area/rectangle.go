package main

import (
	"fmt"
)

type Rectangle struct {
	length, breath float64
}

func (r *Rectangle) Area(){
	area := r.length * r.breath
	fmt.Printf("Area of the Rectangle is: %.2f\n", area)
}