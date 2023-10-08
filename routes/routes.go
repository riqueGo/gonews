package routes

import (
	"net/http"

	"github.com/MrHenri/gonews/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/create-news", controllers.NewNews)
	http.HandleFunc("/form-news", controllers.Form)
}
