package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlerIndex)

	http.HandleFunc("/about", handlerAbout)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe("localhost:9000", nil)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var data = M{"name": "Batman"}
	var tmpl = template.Must(template.ParseFiles(
		"views/index.html",
		"views/_header.html",
		"views/_message.html",
	))

	var err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func handlerAbout(w http.ResponseWriter, r *http.Request) {
	var data = M{"name": "Batman"}
	var tmpl = template.Must(template.ParseFiles(
		"views/about.html",
		"views/_header.html",
		"views/_message.html",
	))

	var err = tmpl.ExecuteTemplate(w, "about", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
