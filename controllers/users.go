package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Error parsing form value", http.StatusBadRequest)
	// 	return
	// }
	// fmt.Fprint(w, "Email: ", r.PostFormValue("email"))
	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
	// fmt.Fprint(w, "Password: ", r.PostFormValue("password"))
}

func RedirectToSignUp(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/users/new", http.StatusSeeOther)
}
