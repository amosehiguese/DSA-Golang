package main

import (
	"fmt"
)

var Pl = fmt.Println

func main(){
	c := CurvedAreaCylinder{10,5}
	c.CurvedArea()
}