package Ascii

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type output struct {
	message string
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Check it again
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "static/500.html")
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

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "static/405.html")
		return
	}

	//Data needed for the templates
	result := ""
	text := ""
	font := ""
	text = r.FormValue("text")
	font = r.FormValue("font")
	if text != "" && CheckAlphabets(text) && font != "" && ValidateFont(font) {
		// 'text' form value is submitted, handle it

		result = FinalResult(text, font)
		// Use 'result' as needed
	} else {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "static/400.html")
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

// this to export a file
func exportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "static/405.html")
		return
	}
	//  Content-Type header to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Content-Disposition header to specify the filename
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")

	//  Content-Length header to the size of the file
	w.Header().Set("Content-Length", strconv.Itoa(len(r.FormValue("ascii"))))

	// Write the ASCII art content to the response writer
	w.Write([]byte(r.FormValue("ascii")))
}

func StartServer() {
	// check it again to understand
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// This for the GET the server directs to the index to html
	http.HandleFunc("/", GetHandler)
	// This for handling the POST request from the form
	http.HandleFunc("/ascii-art", formHandler)
	//this to export a file
	http.HandleFunc("/export", exportHandler)
	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
