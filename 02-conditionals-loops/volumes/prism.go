package main

import (
	"fmt"
)



type Prism struct {
	base, height float64
}

func (p *Prism) Volume() {
	volume := p.base * p.height
	fmt.Printf("Volume of the Prism is: %.2f\n", volume)
}

