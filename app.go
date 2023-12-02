package main

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("intro.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
		}

		err = templ.ExecuteTemplate(w, "intro.html", nil)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	r.Get("/message", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("message.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
		}

		err = templ.ExecuteTemplate(w, "message.html", nil)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
