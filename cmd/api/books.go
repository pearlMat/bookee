package main

import (
	//"encoding/json"
	"fmt"
	"time"

	"net/http"

	
	"github.com/pearlMat/bookee/internal/data"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new book")
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	book := data.Book{
		ID: id,

		CreatedAt: time.Now(),
		Author:    "James Conn",
		Title:     "Casablanca",
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
