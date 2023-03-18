package main

import (
	"fmt"
)



type Parallelogram struct {
	base, height float64

}

func (p *Parallelogram) Area() {
	area := p.base * p.height
	fmt.Printf("Area of the parallelogram is: %.2f\n", area)
}

