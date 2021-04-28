package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const view string = `<html>
<head>
<title>Template</title>
</head>
<body>
<h1>Hello</h1>
</body>
</html>`

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://localhost:9000/index", http.StatusTemporaryRedirect)
	})
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", mux)
}
