package internal

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Books struct {
	ID        int     `json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}