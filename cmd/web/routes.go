package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/participant/new", app.participantNew)
	mux.HandleFunc("/participant/view", app.participantView)
	mux.HandleFunc("/participant/upgrade", app.participantUpgrade)
	mux.HandleFunc("/participant/delete", app.participantDelete)
	mux.HandleFunc("/participant/all", app.participantsAll)
	return mux
}
