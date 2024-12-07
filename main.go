package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BookHandler struct {
	Books []Book `json:"books"`
}

func NewBookHandler(book []Book) *BookHandler {
	return &BookHandler{
		Books: book,
	}
}

//go:generate mockery --name=GetBook --output=./mocks --outpkg=mocks --case=underscore
type GetBook interface {
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
}

func (b *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(&b.Books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}

func (b *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	idInStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idInStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, book := range b.Books {
		if book.ID == id {
			resp, err := json.Marshal(&book)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(resp)
		} else {
			w.WriteHeader(404)
		}
	}
}

func main() {
	book := []Book{
		{
			1,
			"Кафе на краю берега",
		},
		{
			2,
			"Возвращение в кафе",
		},
	}

	bookHandler := NewBookHandler(book)

	r := chi.NewRouter()

	r.Get("/", bookHandler.GetAllBooks)
	r.Get("/{id}", bookHandler.GetBook)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
