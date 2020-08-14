package api

import (
	"net/http"

	"github.com/go-chi/render"

	"eight/internal/domain/authors"
	"eight/internal/domain/books"
)

// API holds all the dependencies required to expose APIs. And each API is a function with *API as its receiver
type API struct {
	books   *books.HandlerBooks
	authors *authors.HandlerAuthors
}

// add microservice to the PARAM
func NewService(bs *books.HandlerBooks, as *authors.HandlerAuthors) (*API, error) {
	return &API{
		books:   bs,
		authors: as,
	}, nil
}

func (a API) HandleLive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("."))
	}
}

func (a API) HandleReady() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.books.Ping()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
	}
}