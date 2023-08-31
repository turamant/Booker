package main

import "net/http"

func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/about", app.aboutHandler)
	mux.HandleFunc("/create", app.createHandler)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}