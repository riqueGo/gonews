package main

import (
	"fmt"
	"net/http"

	"github.com/MrHenri/gonews/db/migration"
	"github.com/MrHenri/gonews/routes"
)

func main() {
	migration.AutoMigration()
	routes.LoadRoutes()

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
