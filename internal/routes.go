package internal

import "github.com/gofiber/fiber/v2"

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", r.CreateBook)
	api.Delete("/books/:id", r.DeleteBook)
	api.Get("/books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
}
