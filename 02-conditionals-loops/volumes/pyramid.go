package main

import (
	"fmt"
)



type Pyramid struct {
	base, height float64
}

func (p *Pyramid) Volume() {
	volume := 0.33 * p.base * p.height
	fmt.Printf("Volume of the Pyramid is: %.2f\n", volume)
}

