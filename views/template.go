package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	tpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, pattern string) (Template, error) {
	tpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		// Failing to parse due to invalid information inside the template.
		return Template{}, fmt.Errorf("FS parsing template: %w", err)
	}

	t := Template{
		tpl: tpl,
	}

	return t, nil
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		// Failing to parse due to invalid information inside the template.
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	t := Template{
		tpl: tpl,
	}

	return t, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.tpl.Execute(w, nil)
	if err != nil {
		// Invalid information during randering
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error executing the templete", http.StatusInternalServerError)
		return
	}
}
