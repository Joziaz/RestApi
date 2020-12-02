package main

const GET string = "GET"
const POST string = "POST"
const PUT string = "PUT"
const DELETE string = "DELETE"

type Article struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Desciption string `json:"descip"`
}

var Articles []Article
