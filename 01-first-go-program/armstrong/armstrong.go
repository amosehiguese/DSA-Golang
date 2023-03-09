package main

import (
	"fmt"
	"math"
)

func main(){
	var firstNum, secondNum, original, sum int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&firstNum)
	
	fmt.Print("Enter second number: ")
	fmt.Scanln(&secondNum)

	for i:=firstNum; i < secondNum; i++{
		original = i
		sum = 0

		for original > 0 {
			rem := original % 10
			sum += int(math.Pow(float64(rem), 3))
			original /= 10
		}

		if sum == i {
			fmt.Println(sum)
		}
	}
}