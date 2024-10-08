package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in go lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move two spots")
	case 3:
		fmt.Println("You can move three spots")
	case 4:
		fmt.Println("You can move four spots")
	case 5:
		fmt.Println("You can move five spots")
	case 6:
		fmt.Println("You can move six spots and roll dice again")
	default:
		fmt.Println("What was that!")
	}

	// there are no break statements in go lang
	// though we can use fallthrough keyword to fall through a case to another one

	/* switch diceNumber {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
		fallthrough
	case 4:
		fmt.Println("Four")
		fallthrough
	case 5:
		fmt.Println("Five")
		fallthrough
	case 6:
		fmt.Println("Six")
	default:
		fmt.Println("What was that!")
	} */
}
