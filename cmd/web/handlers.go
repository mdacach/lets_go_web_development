package main

import (
	"fmt"
	"html/template"
	"net/http"
	"snippetbox/pkg/models"
	"strconv"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	// As "/" acts as a catch-all, we need to actively check
	// to see if the path matched exactly.
	if r.URL.Path != "/" {
		app.notFoundError(w)
		return
	}

	filePaths := []string{
		"./ui/html/home.page.tmpl.html",
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
	}

	ts, err := template.ParseFiles(filePaths...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	idFromQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromQuery)
	// IDs start at 1.
	if err != nil || id < 1 {
		app.notFoundError(w)
		return
	}

	snippet, err := app.snippetModel.Get(id)
	if err == models.ErrorNoRecord {
		app.notFoundError(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%v", snippet)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := "7"

	id, err := app.snippetModel.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
