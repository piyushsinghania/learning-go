package main

import "fmt"

const LoginToken string = "qolsjlvouae" // public

func main() {
	var username string = "piyush"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float32 = 255.345457455
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var intVar int
	fmt.Println(intVar)
	fmt.Printf("Variable is of type: %T \n", intVar)

	var strVar string
	fmt.Println(strVar)
	fmt.Printf("Hey there %s piyush", strVar)
	fmt.Printf("Variable is of type: %T \n", strVar)

	// implicity type
	var college = "Silicon Insititue of Technology"
	fmt.Println(college)

	var age = 23
	fmt.Println(age)

	// no var style
	time := 3.07
	fmt.Println(time)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
