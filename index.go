package main

import "fmt"
import "net/http"
import "html/template"

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server start")
	http.HandleFunc("/", home)
	http.ListenAndServe(":9001", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.New("index").ParseFiles("template/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}