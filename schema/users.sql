-- name: CreateUser :exec
-- Creates a new user
INSERT INTO users (name, email, password, type, sport_center_id, deleted)
VALUES ($1, $2, $3, $4, $5, false);

-- name: GetUser :one
-- Retrieves a user by ID
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
-- Lists all users
SELECT * FROM users;

-- name: ListUserByType :many
-- Lists all users with a given type
SELECT * FROM users WHERE type = $1;

-- name: FilterUsers :many
-- Filters users based on optional parameters (user type, sport center id, deleted)
SELECT * FROM users
WHERE
  ($1::user_type[] IS NULL OR type = ANY($1)) AND
  ($2::integer[] IS NULL OR sport_center_id = ANY($2)) AND
  ($3::boolean IS NULL OR deleted = $3);

-- name: UpdateUser :one
-- Updates a user's information
UPDATE users SET
  name = COALESCE($2, name),
  email = COALESCE($3, email),
  password = COALESCE($4, password),
  type = COALESCE($5, type),
  sport_center_id = COALESCE($6, sport_center_id)
WHERE id = $1
RETURNING *;

-- name: SoftDeleteUser :exec
-- Soft-deletes a user by setting its deleted tag to true
UPDATE users SET deleted = true WHERE id = $1 AND deleted = false RETURNING *;

-- name: ConfirmDeleteUser :exec
-- Deletes a user
DELETE FROM users WHERE id = $1 AND deleted = true;
