package main

import (
	"fmt"
	"net/http"

	"github.com/example.com/app/controllers"
	"github.com/example.com/app/templates"
	"github.com/example.com/app/views"
	"github.com/go-chi/chi/v5"
)

func main () {
	r := chi.NewRouter()

	tpl, err := views.ParseFS(templates.FS, "layout.gohtml", "home.gohtml", "layout-parts.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "layout.gohtml", "contact.gohtml", "layout-parts.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "layout.gohtml", "faq.gohtml", "layout-parts.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.FAQ(tpl))

	tpl, err = views.ParseFS(templates.FS, "layout.gohtml", "pricing.gohtml", "layout-parts.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/pricing", controllers.StaticHandler(tpl))

	usersController := controllers.Users{}
	usersController.Templates.New, err = views.ParseFS(
		templates.FS,
		"auth-layout.gohtml", "signup.gohtml", "layout-parts.gohtml",
	)
	if err != nil {
		panic(err)
	}
	r.Get("/auth/signup", usersController.New)
	r.Post("/users", usersController.Create)

	tpl, err = views.ParseFS(templates.FS, "auth-layout.gohtml", "login.gohtml", "layout-parts.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/auth/login", controllers.StaticHandler(tpl))

	fmt.Printf("Starting server on port 4343")
	err = http.ListenAndServe(":4343", r)
	if err != nil {
		panic(err)
	}

}
