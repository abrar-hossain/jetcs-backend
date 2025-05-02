-- name: CreateReview :one
INSERT INTO reviews (
  article_id, reviewer_id, comments,
  feedback_file_path, decision, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetReviewsForArticle :many
SELECT * FROM reviews WHERE article_id = $1;

-- name: GetReviewsByReviewer :many
SELECT * FROM reviews WHERE reviewer_id = $1;

-- name: CountDecisions :one
SELECT decision, COUNT(*) FROM reviews
WHERE article_id = $1 GROUP BY decision;
