package main

import "fmt"

func main() {
	greet()
	fmt.Println("Functions in go lang")

	result := sum(2, 3)
	fmt.Println("Sum:", result)

	result, ok := dynamicSum(2, 8, 4, 2, 10)
	if ok {
		fmt.Println("Dynamic sum:", result)
	}
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func dynamicSum(values ...int) (int, bool) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, true
}

func greet() {
	fmt.Println("Hola go lang")
}
