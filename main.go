package main

import (
	functions "Ascii/functions"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type output struct {
	message string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	//Data needed for the templates
	text := r.FormValue("text")
	fmt.Println(text)
	font := r.FormValue("font")
	result := functions.FinalResult(text, font)

	if r.Method == "GET" {
		// Parse the template file
		tmpl, err := template.ParseFiles("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Create a data structure to pass to the template
		data := output{
			message: result,
		}

		// Execute the template with the data
		err = tmpl.Execute(w, data.message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err%v", err)
			return
		}

		// Parse the template file
		tmpl, err := template.ParseFiles("static/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a data structure to pass to the template
		data := output{
			message: result,
		}

		// Execute the template with the data
		err = tmpl.Execute(w, data.message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	// This for the GET the server directs to the index to html
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	//This for handling the POST request from the form
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	//fmt.Print(functions.FinalResult())
}
