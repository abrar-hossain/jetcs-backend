-- name: CreateArticle :one
INSERT INTO articles (
  title, abstract, file_path, subject,
  status, author_id, issue_id, created_at
) VALUES (
  $1, $2, $3, $4,
  $5, $6, $7, $8
) RETURNING *;

-- name: GetArticleByID :one
SELECT * FROM articles WHERE id = $1;

-- name: GetArticlesByAuthor :many
SELECT * FROM articles WHERE author_id = $1 ORDER BY created_at DESC;

-- name: ListSubmittedArticles :many
SELECT * FROM articles WHERE status = 'submitted' ORDER BY created_at DESC;
