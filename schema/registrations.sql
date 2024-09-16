-- name: CreateRegistration :one
-- Creates a new registration
INSERT INTO registrations (user_id, course_offering_id, registration_date, price_paid)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetRegistration :one
-- Retrieves a registration by user ID and course offering ID
SELECT * FROM registrations WHERE user_id = $1 AND course_offering_id = $2;

-- name: ListUserRegistrations :many
-- Lists all registrations for a user
SELECT * FROM registrations WHERE user_id = $1;

-- name: ListCourseRegistrations :many
-- Lists all registratiosn for a course
SELECT * FROM registrations WHERE course_offering_id = $1;

-- name: UpdateRegistration :one
-- Updates a registration's information
UPDATE registrations SET registration_date = $3, price_paid = $4
WHERE user_id = $1 AND course_offering_id = $2 RETURNING *;

-- name: DeleteRegistration :exec
-- Deletes a registration
DELETE FROM registrations WHERE user_id = $1 AND course_offering_id = $2;
