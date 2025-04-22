package query

import (
	"database/sql"
	"jetcs-backend/models"
)

func GetAuthorByEmail(db *sql.DB, email string) (*models.Author, error) {
	query := `SELECT id, full_name, email, password, created_at FROM authors WHERE email = $1`

	var author models.Author
	err := db.QueryRow(query, email).Scan(
		&author.ID,
		&author.FullName,
		&author.Email,
		&author.Password,
		&author.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &author, nil
}
