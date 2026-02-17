package internal

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Controller struct {
	repo *Repository
}
func NewController(repo *Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}	

func (c *Controller) CreateBook(ctx *fiber.Ctx) error {
	book := Book{}

	err := ctx.BodyParser(&book)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	id, err := c.repo.CreateBook(book)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Cannot create book",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book created successfully",
		"id":      id,
	})
}

func (c *Controller) DeleteBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	rowsAffected, err := c.repo.DeleteBook(id)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Cannot delete book",
		})
	}

	if rowsAffected == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}

func (c *Controller) GetBooks(ctx *fiber.Ctx) error {
	books, err := c.repo.GetBooks()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Cannot get books",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Books retrieved successfully",
		"data":    books,
	})
}

func (c *Controller) GetBookByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	book, err := c.repo.GetBookByID(id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ctx.Status(404).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"error": "Cannot get book",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book retrieved successfully",
		"data":    book,
	})
}
