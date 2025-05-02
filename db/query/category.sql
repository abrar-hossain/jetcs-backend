-- name: AddUserCategory :exec
INSERT INTO user_categories (user_id, category) VALUES ($1, $2);

-- name: GetUserCategories :many
SELECT category FROM user_categories WHERE user_id = $1;

-- name: FindReviewersByCategory :many
SELECT u.*
FROM users u
JOIN user_categories c ON u.id = c.user_id
WHERE c.category = $1 AND u.user_role = 'reviewer';
