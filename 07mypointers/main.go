package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in golang")

	var ptr *int
	fmt.Println("PTR address is:", ptr)
	// fmt.Println("PTR value is:", *ptr) // since nil pointer we can't do *ptr

	var num int = 23
	var ptr2 *int = &num
	fmt.Println("PTR2 address is:", ptr2)
	fmt.Println("PTR2 value is:", *ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("Mynumber:", *ptr2)
}
