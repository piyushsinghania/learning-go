package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in go")
	content := "Hey, there! How is go lang?"
	filepath := "./note.txt"

	file, err := os.Create(filepath)
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length:", length)
	defer file.Close()

	readFile(filepath)
}

func readFile(filepath string) {
	dataBuffer, err := os.ReadFile(filepath)
	// dataBuffer, err := ioutil.ReadFile(filepath)
	checkNilErr(err)

	fmt.Println(string(dataBuffer))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
