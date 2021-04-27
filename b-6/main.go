package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "Darien Kentanu",
			Gender:  "Male",
			Hobbies: []string{"Music", "Online Shopping", "Coding"},
			Info:    Info{"Darien Holdings, Inc", "New York"},
		}

		tmpl := template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", mux)

}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 division"
}
