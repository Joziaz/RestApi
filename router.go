package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", HandlerHome).Methods(GET)

	r.HandleFunc("/articles", GetAllArticles).Methods(GET)
	r.HandleFunc("/articles", PostArticle).Methods(POST)
	r.HandleFunc("/articles/{id}", GetSingleArticle).Methods(GET)
	r.HandleFunc("/articles/{id}", UpdateArticle).Methods(PUT)
	r.HandleFunc("/articles/{id}", DeleteArticle).Methods(DELETE)

	return r
}
