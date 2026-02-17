package internal

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateBook(book Book) (int, error) {
	query := `INSERT INTO books (author, title, publisher) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.DB.QueryRow(context.Background(), query, book.Author, book.Title, book.Publisher).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) DeleteBook(id string) (int64, error) {
	query := `DELETE FROM books WHERE id = $1`
	commandTag, err := r.DB.Exec(context.Background(), query, id)
	if err != nil {
		return 0, err
	}
	return commandTag.RowsAffected(), nil
}

func (r *Repository) GetBooks() ([]Books, error) {
	query := `SELECT id, author, title, publisher, created_at, updated_at FROM books`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Books
	for rows.Next() {
		var book Books
		err := rows.Scan(&book.ID, &book.Author, &book.Title, &book.Publisher, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *Repository) GetBookByID(id string) (*Books, error) {
	query := `SELECT id, author, title, publisher, created_at, updated_at FROM books WHERE id = $1`
	var book Books
	err := r.DB.QueryRow(context.Background(), query, id).Scan(&book.ID, &book.Author, &book.Title, &book.Publisher, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
