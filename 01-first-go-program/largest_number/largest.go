package main

import (
	"fmt"
)

func main(){
	var x, y int
	fmt.Print("Enter first Number: ")
	fmt.Scanln(&x)
	fmt.Print("Enter second Number: ")
	fmt.Scanln(&y)

	if  x > y {
		fmt.Printf("%d is larger", x)
	}else{
		fmt.Printf("%d is larger", y)
	}
}