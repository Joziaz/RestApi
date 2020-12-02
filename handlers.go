package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "holaa")
}

func GetAllArticles(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
	err := encoder.Encode(Articles)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func GetSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {

		if article.Id == key {

			json.NewEncoder(w).Encode(article)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func PostArticle(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(data, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(data, &article)

	for _, art := range Articles {

		if art.Id == article.Id {
			art = article
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)

}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for index, article := range Articles {
		if article.Id == id {

			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}
