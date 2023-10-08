package controllers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/MrHenri/gonews/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Index : index
func Home(w http.ResponseWriter, r *http.Request) {
	news := models.News{}

	newsList, err := news.GetAllNews()
	if err != nil {
		log.Fatal(err)
	}

	temp.ExecuteTemplate(w, "Index", newsList)
}

// New : new
func NewNews(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		news := models.News{
			Title:     r.FormValue("title"),
			Content:   r.FormValue("content"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		_, err := news.AddNews()
		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

// Form : form
func Form(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Form", nil)
}
