package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")

	user := User{"Piyush", "piyush@gmail.com", 12, true}
	fmt.Println(user)
	fmt.Printf("User detaiils: %+v\n", user)
	fmt.Printf("User age: %v\n", user.Age)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
