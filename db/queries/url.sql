-- name: CreateURL :one
INSERT INTO urls (original_url) VALUES ($1)
RETURNING *;

-- name: GetURLByID :one
SELECT * FROM urls WHERE id = $1;

-- name: GetURLByShortCode :one
SELECT urls.* FROM urls
JOIN redirects ON urls.id = redirects.url_id
WHERE redirects.short_code = $1;
