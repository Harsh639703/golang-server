package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err: %v ", err)
	}
	fmt.Fprintf(w, "POST request successfull\n")
	

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)
	
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)  //this will check the path if it other tha "/hello " then that give 404 not found error  
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)  // if we use the method other than the post then it will give the method not supported error
		return
	}
	fmt.Fprintf(w, "hello!")

}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	// here we created three routes for our server

	http.Handle("/", fileserver)   
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
