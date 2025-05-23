-- name: CreateURL :one
INSERT INTO urls (original_url, short_code)
VALUES ($1, $2)
RETURNING *;

-- name: GetURLByID :one
SELECT * FROM urls WHERE id = $1;

-- name: GetURLByShortCode :one
SELECT * FROM urls WHERE short_code = $1;

-- name: GetAllURLs :many
SELECT * FROM urls;

-- name: UpdateURLClickCount :exec
UPDATE urls SET click_count = click_count + 1 WHERE id = $1;

-- name: UpdateShortCode :exec
UPDATE urls SET short_code = $1 WHERE id = $2;

-- name: DeleteURL :exec
DELETE FROM urls WHERE id = $1;