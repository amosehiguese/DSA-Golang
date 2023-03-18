package main

import (
	"fmt"
)

var Pl = fmt.Println

func main(){
	e := EquiTriangle{5}
	e.Perimeter()
}