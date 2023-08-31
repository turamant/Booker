package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct{
	errorLog *log.Logger
	infoLog  *log.Logger

}


 
func main(){
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "EROR\t", log.Ldate|log.Ltime)
	
	app := application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/about", app.aboutHandler)
	mux.HandleFunc("/create", app.createHandler)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static",fileServer))

    srv := &http.Server{
		Addr : *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Server start on port %s\n", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
	
}