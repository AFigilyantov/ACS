package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)

	}

}

func (app *application) catalog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/catalog/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/catalog.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)

	}

}

// SPORTSMAN HANDLER FUNCTIONS
func (app *application) sportsmanNew(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Add new patricipant"))
}

func (app *application) sportsmanUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change patricipant's data"))
}

func (app *application) spotsmanDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "participant with id %d was deleted ", id)

}
func (app *application) sportsmanView(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "View details of patricipant with id %d", id)

}

func (app *application) sportsmanAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("View full list of patricipants"))
}

// REFEREE HANDLER FUNCTIONS
func (app *application) refereeNew(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Add new REFEREE"))
}

func (app *application) refereeUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change REFEREE data"))
}

func (app *application) refereeDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "REFEREE with id %d was deleted ", id)

}
func (app *application) refereeView(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "View details of REFEREE with id %d", id)

}

func (app *application) refereeAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("View full list of REFEREE"))
}

// CATEGORY HANDLER FUNCTIONS
func (app *application) categoryNew(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Add new category"))
}

func (app *application) categoryUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change category data"))
}

func (app *application) categoryDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "category with id %d was deleted ", id)

}
func (app *application) categoryView(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "View details of category with id %d", id)

}

func (app *application) categoryAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("View full list of category"))
}

// REGION HANDLER FUNCTIONS
func (app *application) regionNew(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Add new REGION"))
}

func (app *application) regionUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change REGION data"))
}

func (app *application) regionDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Region with id %d was deleted ", id)

}
func (app *application) regionView(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "View details of Region with id %d", id)

}

func (app *application) regionAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("View full list of REGIONS"))
}

// SPORTGRADE HANDLER FUNCTIONS
func (app *application) sportGradeNew(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Add new sportGrade"))
}

func (app *application) sportGradeUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change sportGrade data"))
}

func (app *application) sportGradeDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "sportGrade with id %d was deleted ", id)

}
func (app *application) sportGradeView(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "View details of sportGrade with id %d", id)

}

func (app *application) sportGradeAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("View full list of sportGrade"))
}

// REFEREEGRADE HANDLER FUNCTIONS
func (app *application) refereeGradeNew(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Add new refereeGrade"))
}

func (app *application) refereeGradeUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change refereeGrade data"))
}

func (app *application) refereeGradeDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "refereeGrade with id %d was deleted ", id)

}
func (app *application) refereeGradeView(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "View details of refereeGrade with id %d", id)

}

func (app *application) refereeGradeAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("View full list of refereeGrade"))
}

// REFEREEROLE HANDLER FUNCTIONS

func (app *application) refereeRoleNew(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Add new refereeRole"))
}

func (app *application) refereeRoleUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change refereeRole data"))
}

func (app *application) refereeRoleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "refereeRole with id %d was deleted ", id)

}
func (app *application) refereeRoleView(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "View details of refereeRole with id %d", id)

}

func (app *application) refereeRoleAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("View full list of refereeRole"))
}

// GENDER HANDLER FUNCTIONS

func (app *application) genderNew(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Add new gender"))
}

func (app *application) genderUpgrade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.Header().Set("Allow", http.MethodPut)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change gender data"))
}

func (app *application) genderDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "gender with id %d was deleted ", id)

}
func (app *application) genderView(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "View details of gender with id %d", id)

}

func (app *application) genderAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("View full list of gender"))
}
