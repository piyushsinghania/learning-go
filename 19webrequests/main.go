package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://formester.com"

func main() {
	fmt.Println("Web requests in go")
	response, err := http.Get(url)
	fmt.Printf("Type of response is: %T\n", response)
	defer response.Body.Close()
	checkNilErr(err)

	dataBuffer, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	var pageContent string = string(dataBuffer)

	fmt.Println(pageContent)

	filepath := "./page.txt"
	file, err := os.Create(filepath)
	checkNilErr(err)

	io.WriteString(file, pageContent)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
