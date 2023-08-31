package main

import (
	"html/template"
	"net/http"
)

func (app *application) homeHandler(responce http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		app.notFound(responce)
		return
	}
	files := []string{
		"./ui/html/index.html",
		"./ui/html/base.html",
		"./ui/html/footer.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(responce, err)
		return
	}
	err = ts.Execute(responce, nil)
	if err != nil {
		app.serverError(responce, err)
	}

}

func (app *application) createHandler(responce http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		responce.Header().Set("allow", http.MethodPost)
		responce.Header().Set("Content-Type", "application/json")
		// responce.WriteHeader(405)
		// responce.Write([]byte("Not Allowed Method"))
		app.clientError(responce, http.StatusMethodNotAllowed)
		return
	}
	responce.Write([]byte("Create cnippet"))

}
func (app *application) aboutHandler(responce http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/about" {
		app.notFound(responce)
		return
	}
	files := []string{
		"./ui/html/about.html",
		"./ui/html/base.html",
		"./ui/html/footer.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(responce, err)
	}
	err = ts.Execute(responce, nil)
	if err != nil {
		app.serverError(responce, err)
	}
}
