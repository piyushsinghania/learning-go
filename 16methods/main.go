package main

import "fmt"

func main() {
	fmt.Println("Methods in go lang")
	user := User{"Piyush", "piyush@gmail.com", 12, true}
	fmt.Println(user)
	fmt.Printf("User detaiils: %+v\n", user)
	fmt.Printf("User age: %v\n", user.Age)

	user.GetStatus()
	user.NewMail()

	fmt.Printf("User detaiils: %+v\n", user)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("New user email:", u.Email)

}
