package internal

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := Book{}

	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var id int
	err = r.DB.QueryRow(context.Background(), CreateBookQuery, book.Author, book.Title, book.Publisher).Scan(&id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Cannot create book",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book created successfully",
		"id":      id,
	})
}

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	commandTag, err := r.DB.Exec(context.Background(), DeleteBookQuery, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Cannot delete book",
		})
	}

	if commandTag.RowsAffected() == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	rows, err := r.DB.Query(context.Background(), GetBooksQuery)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Cannot get books",
		})
	}
	defer rows.Close()

	var books []Books
	for rows.Next() {
		var book Books
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Publisher, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Cannot scan book",
			})
		}
		books = append(books, book)
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Books retrieved successfully",
		"data":    books,
	})
}

func (r *Repository) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Books
	err := r.DB.QueryRow(context.Background(), GetBookByIDQuery, id).Scan(&book.ID, &book.Author, &book.Title, &book.Publisher, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Cannot get book",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book retrieved successfully",
		"data":    book,
	})
}

