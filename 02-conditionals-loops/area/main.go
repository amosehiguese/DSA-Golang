package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	c := Cube{10}
	c.SurfaceArea()
	pl(c)
}