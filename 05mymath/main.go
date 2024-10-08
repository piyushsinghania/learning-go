package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var x int = 2
	// var y float32 = 4.5
	// fmt.Println("The sum is:", x+int(y))

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNumber)
}
