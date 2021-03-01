package main

import (
	"net/http"
	"html/template"
	"path"
	"fmt"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Golang HTML",
			"name":  "Hello World!",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9090")
	http.ListenAndServe(":9090", nil)
}