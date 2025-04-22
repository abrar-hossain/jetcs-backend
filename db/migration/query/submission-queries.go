package query

import (
	"database/sql"
	"jetcs-backend/models"
)

// CreateSubmission inserts a new submission and returns the result
func CreateSubmission(db *sql.DB, submission *models.Submission) error {
	query := `
        INSERT INTO submissions (title, author_name, email)
        VALUES ($1, $2, $3)
        RETURNING id, status
    `

	return db.QueryRow(query, submission.Title, submission.AuthorName, submission.Email).
		Scan(&submission.Id, &submission.Status)
}
