package main

import "fmt"

func main() {
	fmt.Println("If else in go")

	loginCount := 8
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount == 10 {
		result = "Active user"
	} else {
		result = "Suspicious user"
	}

	fmt.Println(result)

	if x := 3; x%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
