package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, value := range days {
		fmt.Printf("Index: %v => Value: %v\n", index, value)
	}

	j := 1

	for j < 10 {
		if j == 5 {
			j++
			// break
			continue
		} else if j == 7 {
			goto portfolio
		}

		fmt.Println(j)
		j++
	}

portfolio:
	fmt.Println("piyushsinghania.com")
}
