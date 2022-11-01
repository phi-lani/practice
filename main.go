package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/phi-lani/webApp/controllers"
	"github.com/phi-lani/webApp/templates"
	"github.com/phi-lani/webApp/views"
)

func main() {

	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "about.gohtml", "tailwind.gohtml"))
	r.Get("/about", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
