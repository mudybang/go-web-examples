package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
	}
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Bagoes"}
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Bagoes"}
		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:3000")
	http.ListenAndServe("localhost:9000", nil)
}
