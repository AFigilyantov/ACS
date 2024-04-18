package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	// GETTING STATIC CONTENT
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//HANDLER REGISTRATION
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("GET /catalog/", app.catalog)

	// SPORTSMEN ROUTES
	mux.HandleFunc("POST /sportsman/new", app.sportsmanNew)
	mux.HandleFunc("GET /sportsman/{id}/view", app.sportsmanView)
	mux.HandleFunc("PUT /sportsman/upgrade", app.sportsmanUpgrade)
	mux.HandleFunc("DELETE /sportsman/{id}/delete", app.spotsmanDelete)
	mux.HandleFunc("GET /sportsman/all", app.sportsmanAll)

	// REFEREE ROUTES
	mux.HandleFunc("POST /referee/new", app.refereeNew)
	mux.HandleFunc("GET /referee/{id}/view", app.refereeView)
	mux.HandleFunc("PUT /referee/upgrade", app.refereeUpgrade)
	mux.HandleFunc("DELETE /referee/{id}/delete", app.refereeDelete)
	mux.HandleFunc("GET /referee/all", app.refereeAll)

	// CATEGORY ROUTES
	mux.HandleFunc("POST /category/new", app.categoryNew)
	mux.HandleFunc("GET /category/{id}/view", app.categoryView)
	mux.HandleFunc("PUT /category/upgrade", app.categoryUpgrade)
	mux.HandleFunc("DELETE /category/{id}/delete", app.categoryDelete)
	mux.HandleFunc("GET /category/all", app.categoryAll)

	// SPORTGRADE ROUTES
	mux.HandleFunc("POST /sportGrade/new", app.sportGradeNew)
	mux.HandleFunc("GET /sportGrade/{id}/view", app.sportGradeView)
	mux.HandleFunc("PUT /sportGrade/upgrade", app.sportGradeUpgrade)
	mux.HandleFunc("DELETE /sportGrade/{id}/delete", app.sportGradeDelete)
	mux.HandleFunc("GET /sportGrade/all", app.sportGradeAll)

	// REFEREEGRADE ROUTES
	mux.HandleFunc("POST /refereeGrade/new", app.refereeGradeNew)
	mux.HandleFunc("GET /refereeGrade/{id}/view", app.refereeGradeView)
	mux.HandleFunc("PUT /refereeGrade/upgrade", app.refereeGradeUpgrade)
	mux.HandleFunc("DELETE /refereeGrade/{id}/delete", app.refereeGradeDelete)
	mux.HandleFunc("GET /refereeGrade/all", app.refereeGradeAll)

	// REFEREEROLE ROUTES
	mux.HandleFunc("POST /refereerole/new", app.refereeRoleNew)
	mux.HandleFunc("GET /refereerole/{id}/view", app.refereeRoleView)
	mux.HandleFunc("PUT /refereerole/upgrade", app.refereeRoleUpgrade)
	mux.HandleFunc("DELETE /refereerole/{id}/delete", app.refereeRoleDelete)
	mux.HandleFunc("GET /refereerole/all", app.refereeRoleAll)

	// GENDER ROUTES
	mux.HandleFunc("POST /gender/new", app.genderNew)
	mux.HandleFunc("GET /gender/{id}/view", app.genderView)
	mux.HandleFunc("PUT /gender/upgrade", app.genderUpgrade)
	mux.HandleFunc("DELETE /gender/{id}/delete", app.genderDelete)
	mux.HandleFunc("GET /gender/all", app.genderAll)

	// REGION ROUTES
	mux.HandleFunc("POST /region/new", app.regionNew)
	mux.HandleFunc("GET /region/{id}/view", app.regionView)
	mux.HandleFunc("PUT /region/upgrade", app.regionUpgrade)
	mux.HandleFunc("DELETE /region/{id}/delete", app.regionDelete)
	mux.HandleFunc("GET /region/all", app.regionAll)

	return mux
}
