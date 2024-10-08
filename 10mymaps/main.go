package main

import "fmt"

func main() {
	fmt.Println("Maps in Go language")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("Languages:", languages)
	fmt.Println("Js is for:", languages["js"])

	delete(languages, "py")
	fmt.Println("Remaining languages", languages)

	// looping maps
	for key, value := range languages {
		fmt.Printf("%v => %v\n", key, value)
	}
}
