package main

import (
	"fmt"
)

func main(){
	var x, y int
	var op string
	fmt.Print("Enter first number:")
	fmt.Scanln(&x)
	fmt.Print("Enter second number:")
	fmt.Scanln(&y)
	fmt.Print("Enter operator (+, -, /, *):")
	fmt.Scanln(&op)

	var result int
	var err error 
	switch op {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "/":
		result = x / y
	case "*":
		result = x * y
	default:
		err = fmt.Errorf("not a valid operator")
		fmt.Println(err)
	}
	if err == nil{
		fmt.Printf("%d + %d = %v", x,y, result)
	}
}