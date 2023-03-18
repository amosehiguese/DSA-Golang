package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	c := Cone{10,5}
	c.Volume()
	pl(c)
}