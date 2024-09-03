package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func ShowIndex(w http.ResponseWriter, r *http.Request)  {
	products := models.SearchAllProducts()
	templates.ExecuteTemplate(w, "index", products)
}

func New(w http.ResponseWriter, r *http.Request)  {
	templates.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Erro ao obter o preço", err)
	}
	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil {
		log.Println("Erro ao obter a quantidade", err)
	}

	models.InsertProduct(name, description, price, amount)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	product := models.SearchById(id)
	templates.ExecuteTemplate(w, "edit", product)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println("Erro ao obter o id", err)
	}
	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Erro ao obter o preço", err)
	}
	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil {
		log.Println("Erro ao obter a quantidade", err)
	}

	models.UpdateProduct(id, name, description, price, amount)
	http.Redirect(w, r, "/", 301)
}

func Remove(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}
	models.RemoveProduct(id)
	http.Redirect(w, r, "/", 301)
}
