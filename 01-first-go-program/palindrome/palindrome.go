package main

import (
	"fmt"
)

func isPalindrome(w string) bool {
		i := 0
		j := len(w) - 1

	for i < j {
		if w[i] != w[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main(){
	var word string
	fmt.Print("Enter word: ")
	fmt.Scanln(&word)
	
	fmt.Println(isPalindrome(word))

}