package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please rate this course:")

	// comma ok syntax || comma error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating: ", input)
	fmt.Printf("Type of rating: %T \n", input)
}
