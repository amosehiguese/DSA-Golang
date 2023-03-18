package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	p := Prism{10,5}
	p.Volume()
	pl(p)
}