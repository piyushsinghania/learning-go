package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Guava"

	fmt.Println("Fruitlist:", fruitList)
	fmt.Println("Fruitlist length:", len(fruitList)) // 4

	var vegList = [5]string{"potato", "green peas", "cabbage"}
	fmt.Println("Veglist:", vegList)
	fmt.Println("Veglist length:", len(vegList)) // 5
}
