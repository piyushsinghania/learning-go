package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://piyushsinghania.com:3000/contact?type=phone&address=true"

func main() {
	fmt.Println("Url in go lang")
	parsedUrl, err := url.Parse(myUrl)
	checkNilErr(err)

	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Port())
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.RawQuery)
	fmt.Println(parsedUrl.Query())
	fmt.Println(parsedUrl.Query().Get("type"))
	fmt.Println(parsedUrl.Query().Get("address"))

	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "formester.com",
		Path:     "features",
		RawQuery: "mobile=false",
	}

	fmt.Println(newUrl.String())
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
