package main

import (
	"net/http"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.Get("/", a.Handlers.Home)

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.Routes
}
