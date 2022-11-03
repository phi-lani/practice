package controllers

import (
	"fmt"
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

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprint(w, "Email:", r.FormValue("email"))
	fmt.Fprint(w, "Password:", r.FormValue("password"))

}
