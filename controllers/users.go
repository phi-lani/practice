package controllers

import (
	"net/http"

	"github.com/phi-lani/webApp/views"
)

type Users struct {
	Tempalates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Tempalates.New.Execute(w, nil)
}
