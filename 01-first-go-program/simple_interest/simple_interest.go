package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	var p, t, r float32 
	fmt.Print("Enter principal amount: ")
	fmt.Scanln(&p)
	fmt.Print("Enter rate: ")
	fmt.Scanln(&r)
	fmt.Print("Enter time: ")
	fmt.Scanln(&t)

	simpleInterest := p * r * t
	pl("Your Interest :", simpleInterest)

}