package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func jsonResponse(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write([]byte("\n"))
}

func main()  {
	http.HandleFunc("/api/students", func(w http.ResponseWriter, r *http.Request) {
		if !Auth(w, r) {
			return
		}
		
		if id := r.URL.Query().Get("id"); id != "" {
			jsonResponse(w, GetStudent(id))
		} else {
			http.Error(w, "Invalid id", http.StatusBadRequest)
		}
	})

	fmt.Println("server started at localhost:9090")
	http.ListenAndServe(":9090", nil)
}