package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go")

	var fruitList = []string{"Apple", "Guava", "Banana"}
	fmt.Println("Fruitlist:", fruitList)
	fmt.Println("Length:", len(fruitList))

	fruitList = append(fruitList, "Mango", "Pears")
	fmt.Println("Fruitlist:", fruitList)
	fmt.Println("Length:", len(fruitList))

	fruitList = append(fruitList[0:3])
	fmt.Println("Fruitlist:", fruitList)
	fmt.Println("Length:", len(fruitList))

	fmt.Println(append(fruitList[:2]))
	fmt.Println("Fruitlist:", fruitList)
	fmt.Println("Length:", len(fruitList))

	highScores := make([]int, 4)

	highScores[0] = 320
	highScores[1] = 120
	highScores[2] = 420
	highScores[3] = 520

	fmt.Println("Highscores:", highScores)

	highScores = append(highScores, 20, 620, 720)
	fmt.Println("New highscores:", highScores)

	fmt.Println("Are scores sorted:", sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println("Sorted highscores:", highScores)
	fmt.Println("Are scores sorted:", sort.IntsAreSorted(highScores))

	var languages = []string{"Python", "Javascript", "Java", "Ruby", "Go"}
	var index int = 2
	languages = append(languages[:index], languages[(index+1):]...)
	fmt.Println("Filtered languages:", languages)
}
