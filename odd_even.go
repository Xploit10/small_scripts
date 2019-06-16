package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")

	var number int
	fmt.Scanln(&number)

	if (number % 2 == 0) {
		fmt.Println(number, "is an Even number")

	} else {

		fmt.Println(number, "is an Odd number")
	}
}
