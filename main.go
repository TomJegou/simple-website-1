package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/homepage.html"))
		tmpl.Execute(w, nil)
	})
	fmt.Println("http://localhost:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
