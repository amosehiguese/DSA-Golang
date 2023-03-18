package main

import (
	"fmt"
)

var Pl = fmt.Println

func main(){
	i := Isosceles{5, 10}
	i.Area()
}