package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Calculation struct {
	Result int
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		value1, err1 := strconv.Atoi(r.FormValue("value1"))
		value2, err2 := strconv.Atoi(r.FormValue("value2"))
		if err1 == nil && err2 == nil {
			result := value1 + value2
			tmpl, err := template.ParseFiles("templates/index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, Calculation{Result: result})
			return
		}
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", calculateHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}