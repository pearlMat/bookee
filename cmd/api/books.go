package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/pearlMat/bookee/internal/data"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string   `json:"title"`
		Author string   `json:"author"`
		Isbn   int32    `json:"isbn"`
		Year   int32    `json:"year,omitempty"`
		Genres []string `json:"genres,omitempty"`

		Version int32
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
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
