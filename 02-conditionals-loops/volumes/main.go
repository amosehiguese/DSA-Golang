package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	s := Sphere{5}
	s.Volume()
	pl(s)
}