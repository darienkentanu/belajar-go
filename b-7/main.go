package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Orang struct {
	Name    string
	Alias   string
	Friends []string
}

func (s Orang) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Orang{
			Name:    "Darien Kentanu",
			Alias:   "Ocoy",
			Friends: []string{"Angga", "Firman", "Jonathan"},
		}

		tmpl := template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", mux)
}
