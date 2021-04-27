package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Orang struct {
	Name    string
	Alias   string
	Friends []string
}

func (o Orang) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		orang := Orang{
			Name:    "Darien Kentanu",
			Alias:   "D'Fall",
			Friends: []string{"Angga", "Firman", "Jonathan"},
		}

		tmpl := template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, orang); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", mux)
}
