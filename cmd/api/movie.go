package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new Movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(w, r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
