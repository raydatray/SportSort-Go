-- name: CreateUser :one
-- Creates a new user
INSERT INTO users (name, email, password, type, sport_center_id)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUser :one
-- Retrieves a user by ID
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
-- Lists all users
SELECT * FROM users;


-- name: ListUserByType :many
-- Lists all users with a given type
SELECT * FROM users
WHERE type = $1;

-- name: UpdateUser :one
-- Updates a user's information
UPDATE users SET name = $2, email = $3, password = $4, type = $5, sport_center_id = $6
WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
-- Deletes a user
DELETE FROM users WHERE id = $1;
