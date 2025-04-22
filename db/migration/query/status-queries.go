package query

import (
	"database/sql"
	"errors"
)

// CheckSubmissionStatus returns the status of a submission by its ID
func CheckSubmissionStatus(db *sql.DB, id int) (string, error) {
	var status string
	query := `SELECT status FROM submissions WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&status)

	if err == sql.ErrNoRows {
		return "", errors.New("submission not found")
	} else if err != nil {
		return "", err
	}

	return status, nil
}
