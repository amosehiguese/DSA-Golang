package main

import "fmt"

func getFibo() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

func main() {
	var num int
	fmt.Print("Enter length of fibonacci series: ")
	fmt.Scanln(&num)
	fibVal := getFibo()
	for i := 0; i < num; i++{
		fmt.Println(fibVal())
	}
}