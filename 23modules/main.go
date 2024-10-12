package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules in go lang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/contact", serveContact).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world in go lang</h1>"))
}

func serveContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Contact me at piyushsinghania6@gmail.com"))
}
