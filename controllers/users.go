package controllers

import (
	"fmt"
	"net/http"
)

type Users struct{
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// Render the signup page
	u.Templates.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// Create the user
	err := r.ParseForm()
	if err != nil {
		// panic(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Email: %v\n", r.PostForm.Get("email"))
	fmt.Fprintf(w, "Password: %v", r.PostForm.Get("password"))
}