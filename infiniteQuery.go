package main

import (
	"fmt"
)

var userInput int

func main() {
	//Insert your code here

	for true { //not sure how to use preceeding var declaration for while loop

		fmt.Println("Please input a number")
		fmt.Scan(&userInput)

		if userInput%2 == 1 {
			fmt.Println(userInput, " is an odd number")
		} else {
			fmt.Println(userInput, " is an even number")
		}

		if userInput/10 >= 1 {
			fmt.Println(userInput, " is a double digit number")
		} else {
			fmt.Println(userInput, " is a single digit number")
		}
	}
}
