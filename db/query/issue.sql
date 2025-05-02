-- name: CreateIssue :one
INSERT INTO issues (
  volume, issue_number, publish_date
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetLatestIssue :one
SELECT * FROM issues ORDER BY publish_date DESC LIMIT 1;

-- name: GetIssueByID :one
SELECT * FROM issues WHERE id = $1;

-- name: ListIssues :many
SELECT * FROM issues ORDER BY publish_date DESC;
