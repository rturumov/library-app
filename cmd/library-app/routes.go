package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/books", app.listBooksHandler)
	router.HandlerFunc(http.MethodPost, "/v1/books", app.createBookHandler)
	router.HandlerFunc(http.MethodGet, "/v1/books/:id", app.showBookHandler)
	router.HandlerFunc(http.MethodPut, "/v1/books/:id", app.updateBookHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/books/:id", app.deleteBookHandler)

	router.HandlerFunc(http.MethodGet, "/v1/manga", app.listMangasHandler)
	router.HandlerFunc(http.MethodPost, "/v1/manga", app.createMangaHandler)
	router.HandlerFunc(http.MethodGet, "/v1/manga/:id", app.showMangaHandler)
	router.HandlerFunc(http.MethodPut, "/v1/manga/:id", app.updateMangaHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/manga/:id", app.deleteMangaHandler)

	router.HandlerFunc(http.MethodGet, "/v1/authors", app.listAuthorsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/authors", app.createAuthorHandler)
	router.HandlerFunc(http.MethodGet, "/v1/authors/:id", app.showAuthorHandler)
	router.HandlerFunc(http.MethodPut, "/v1/authors/:id", app.updateAuthorHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/authors/:id", app.deleteAuthorHandler)
	router.HandlerFunc(http.MethodGet, "/v1/authors/:id/books", app.listBooksByAuthorHandler)
	router.HandlerFunc(http.MethodGet, "/v1/authors/:id/manga", app.listMangaByAuthorHandler)

	return app.recoverPanic(app.rateLimit(router))
}
