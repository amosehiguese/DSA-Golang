package main

import (
	"fmt"
)

var Pl = fmt.Println

func main(){
	c := Cube{10}
	c.SurfaceArea()
}