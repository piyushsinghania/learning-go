package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	deferedSlice()
}

func deferedSlice() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// predicted output
// Hello, 4, 3, 2, 1, 0, World
