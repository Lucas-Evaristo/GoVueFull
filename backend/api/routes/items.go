package routes

import (
	"github.com/lucas-evaristo/govuefull/backend/api/handlers"

	"github.com/go-chi/chi"
)

// Registers routes for items
func RegisterItemsRoutes(r chi.Router) {
	r.Get("/items", handlers.GetItems)
	r.Get("/items/{id}", handlers.GetItemByID)
	r.Post("/items", handlers.CreateItem)
	r.Put("/items/{id}", handlers.UpdateItem)
	r.Delete("/items/{id}", handlers.DeleteItem)
}
