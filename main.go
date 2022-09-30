package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Welcome to my great site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1> <p>To get in touch, email me at <a href=\"mailto: philaningumede@gmail.com\">philaningumede@gmail.com</a>.")
}

func fqaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,
		"<h1>FQA</h1><p>Q: Is there a free version?</p><p>A: Yes we offer a free 30 days trial</p>")
}
func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/fqa":
		fqaHandler(w, r)
	default:
		http.NotFound(w, r)
	}

}

func main() {
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
