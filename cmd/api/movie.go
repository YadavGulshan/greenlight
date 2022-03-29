package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/yadavgulshan/greenlight/internal/data"
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

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Barfi",
		Year:      2020,
		Runtime:   102,
		Genre:     []string{"comedy", "action", "drama"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "We're unable to serve your request", http.StatusInternalServerError)
	}
}
