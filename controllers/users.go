package controllers

import (
	"net/http"
)

type Users struct {
	Tempalates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Tempalates.New.Execute(w, nil)
}
