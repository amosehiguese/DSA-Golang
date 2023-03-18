package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	p := Cylinder{8,5}
	p.Volume()
	pl(p)
}