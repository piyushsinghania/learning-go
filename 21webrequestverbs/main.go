package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const getUrl string = "http://localhost:8000/get"
const postUrl string = "http://localhost:8000/post"
const postFormUrl string = "http://localhost:8000/postform"

func main() {
	fmt.Println("Web request methods in go")
	PerformGetRequest(getUrl)
	PerformPostRequest(postUrl)
	PerformPostFromRequest(postFormUrl)
}

func PerformGetRequest(getUrl string) {
	response, err := http.Get(getUrl)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println(string(content))

	// var responseBuilder strings.Builder
	// responseBuilder.Write(content)
	// fmt.Println(responseBuilder.String())
}

func PerformPostRequest(postUrl string) {
	requestBody := strings.NewReader(`
		{
			"name": "piyush singhania",
			"age": 16,
			"college": "Silicon institute of technology",
			"email": "piyush@gmail.com"
		}
	`)

	response, err := http.Post(postUrl, "application/json", requestBody)
	checkNilErr(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println(string(content))
}

func PerformPostFromRequest(postFormUrl string) {
	data := url.Values{}
	data.Add("firstname", "piyush")
	data.Add("lastname", "singhania")
	data.Add("email", "piyush@gmail.com")

	response, err := http.PostForm(postFormUrl, data)
	checkNilErr(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println(string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
