package routes

import (
	"net/http"

	"github.com/LucasEvaristo/GoVueFull/api/handlers"
	"github.com/go-chi/chi"
)

// RegisterRoutes registers all the API routes
func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/items", handlers.GetUsers)
	r.Get("/items/{id}", handlers.GetUserByID)
	r.Post("/items", handlers.CreateUser)
	r.Put("/items/{id}", handlers.UpdateUser)
	r.Delete("/items/{id}", handlers.DeleteUser)

	return r
}
