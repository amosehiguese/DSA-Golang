package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	p := Pyramid{5,8}
	p.Volume()
	pl(p)
}