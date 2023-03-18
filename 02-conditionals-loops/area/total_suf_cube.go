package main

import (
	"fmt"
	"math"
)

type Cube struct {
	sides float64
}

func (c *Cube) SurfaceArea() {
	surfaceAreaOfCube := 6 * math.Pow(c.sides, 2)
	fmt.Printf("The total surface area of a cube is: \n", surfaceAreaOfCube)
}