package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/direct", func(w http.ResponseWriter, r *http.Request) {
		var message = "Langsung Jaya"
		w.Write([]byte(message))
	})
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}
