package main

import (
	"fmt"
) 

func main() {
			var num int;
	 		_, err := fmt.Scanln(&num)
    if err != nil {
					fmt.Println(err)
				}
				if num % 2 == 0 {
					fmt.Println("even")
				} else {
					fmt.Println("odd")
				}
}