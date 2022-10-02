package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/phi-lani/webApp/controllers"
	"github.com/phi-lani/webApp/views"
)

// func executeTemplate(w http.ResponseWriter, filepath string) {
// 	t, err := views.Parse(filepath)

// 	if err != nil {
// 		// Invalid information during randering
// 		log.Printf("Parsing template: %v", err)
// 		http.Error(w, "There was an error executing the templete", http.StatusInternalServerError)
// 		return
// 	}

// 	t.Execute(w, nil)
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	executeTemplate(w, "templates/home.gohtml")
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {
// 	executeTemplate(w, "templates/contact.gohtml")
// }

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// 	executeTemplate(w, "templates/faq.gohtml")
// }

// func notFound(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, "Page Not Found", http.StatusNotFound)
// }

func main() {

	r := chi.NewRouter()

	tpl, err := views.Parse("templates/home.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/contact.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/faq.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))
	// r.Get("/", homeHandler)
	// r.Get("/contact", contactHandler)
	// r.Get("/faq", faqHandler)
	// r.NotFound(notFound)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
