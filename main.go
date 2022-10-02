package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		// Failing to parse due to invalid information inside the template.
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the templete", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		// Invalid information during randering
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error executing the templete", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1> <p>To get in touch, email me at <a href=\"mailto: philaningumede@gmail.com\">philaningumede@gmail.com</a>.")
}

func fqaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FQA</h1><p>Q: Is there a free version?</p><p>A: Yes we offer a free 30 days trial</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", fqaHandler)
	r.NotFound(notFound)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
