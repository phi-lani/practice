package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/phi-lani/webApp/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)

	if err != nil {
		// Invalid information during randering
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error executing the templete", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/contact.gohtml")
}

func fqaHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/faq.gohtml")
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
