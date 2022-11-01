package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

// Template struct from the HTML template package.
type Template struct {
	tpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		// Failing to parse due to invalid information inside the template.
		return Template{}, fmt.Errorf("FS parsing template: %w", err)
	}

	t := Template{
		tpl: tpl,
	}

	return t, nil
}

// Parsing the file before we execute it.
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

// Assuming that at this point the template is already parsed and ready
// to be executed.
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.tpl.Execute(w, data)
	if err != nil {
		// Invalid information during randering
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error executing the templete", http.StatusInternalServerError)
		return
	}
}
