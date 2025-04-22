package query

import (
	"database/sql"
	"jetcs-backend/models"
)

// Save inserts a new author into the PostgreSQL database
func SaveAuthor(db *sql.DB, author *models.Author) error {
	query := `
        INSERT INTO authors (full_name, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `

	return db.QueryRow(query, author.FullName, author.Email, author.Password).
		Scan(&author.ID, &author.CreatedAt)
}
