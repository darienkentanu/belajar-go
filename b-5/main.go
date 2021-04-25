package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	// tmpl, err := template.ParseGlob("views/*")
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }
	// ...
	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(rw http.ResponseWriter, r *http.Request) {
		data := M{"name": "Darien"}
		// err = tmpl.ExecuteTemplate(rw, "index", data)
		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(rw, "index", data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/about", func(rw http.ResponseWriter, r *http.Request) {
		data := M{"name": "Kentanu"}
		// err = tmpl.ExecuteTemplate(rw, "about", data)
		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(rw, "about", data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", mux)

}
