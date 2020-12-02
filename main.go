package main

import (
	"net/http"
)

func main() {

	Articles = []Article{
		{
			Id:         "1",
			Title:      "Tituloooo perros",
			Desciption: "blablablabla, weeeeeo esto es un descrip",
		},
		{
			Id:         "2",
			Title:      "Tituloooo 2",
			Desciption: "weeeeeo esto es un descrip 2",
		},
	}

	http.ListenAndServe(":8000", Router())
}
