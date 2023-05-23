package main

import (
	"net/http"

	"github.com/lucas-evaristo/govuefull/backend/api/routes"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	routes.RegisterItemsRoutes(router)
	http.ListenAndServe(":3000", router)
}
