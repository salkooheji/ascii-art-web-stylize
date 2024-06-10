package Ascii

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type output struct {
	message string
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Check it again
	if r.URL.Path != "/" || r.URL.Path == "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
		return
	}
	// Parse the template file
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
		return
	}
	// Create a data structure to pass to the template
	data := output{
		message: "",
	}
	// Execute the template with the data
	err = tmpl.Execute(w, data.message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	//Data needed for the templates
	result := ""
	text := ""
	font := ""
	if text = r.FormValue("text"); text != "" {
		// 'text' form value is submitted, handle it
		font = r.FormValue("font")
		result = FinalResult(text, font)
		// Use 'result' as needed
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
		return
	}

	//What is this ??
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

func StartServer() {
	// check it again to understand
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// This for the GET the server directs to the index to html
	http.HandleFunc("/", GetHandler)
	// This for handling the POST request from the form
	http.HandleFunc("/ascii-art", formHandler)
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
