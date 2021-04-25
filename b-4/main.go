package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := path.Join("views", "index.html")
		tmpl := template.Must(template.ParseFiles(filepath))
		data := map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Batman",
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", mux)
}
