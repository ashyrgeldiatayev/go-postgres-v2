package internal

import "github.com/gofiber/fiber/v2"

func (c *Controller) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", c.CreateBook)
	api.Delete("/books/:id", c.DeleteBook)
	api.Get("/books/:id", c.GetBookByID)
	api.Get("/books", c.GetBooks)
}
