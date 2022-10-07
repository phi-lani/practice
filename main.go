package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/phi-lani/webApp/controllers"
	"github.com/phi-lani/webApp/views"
)

func main() {

	r := chi.NewRouter()

	tpl := views.Must(views.Parse("templates/home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse("templates/contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse("templates/faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
