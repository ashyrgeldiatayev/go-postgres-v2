package internal

const (
	CreateBookQuery = `INSERT INTO books (author, title, publisher) VALUES ($1, $2, $3) RETURNING id`
	
	DeleteBookQuery = `DELETE FROM books WHERE id = $1`
	
	GetBooksQuery = `SELECT id, author, title, publisher, created_at, updated_at FROM books`
	
	GetBookByIDQuery = `SELECT id, author, title, publisher, created_at, updated_at FROM books WHERE id = $1`
	
	//migration 
	CreateBooksTableQuery = `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		author VARCHAR(255),
		title VARCHAR(255),
		publisher VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
)
