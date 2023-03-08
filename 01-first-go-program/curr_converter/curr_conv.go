package main

import (
	"fmt"
)

func main(){
	var rupCurr, usdCurr int
	fmt.Print("Enter currency in rupees: ")
	fmt.Scanln(&rupCurr)

	usdCurr = rupCurr / 75

	fmt.Printf("%d rupees converts to %d usd", rupCurr, usdCurr)
}