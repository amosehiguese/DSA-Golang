package main

import (
	"fmt"
)

var pl = fmt.Println


func main(){
	_ = pl
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello %s", name)
}