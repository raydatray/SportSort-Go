-- name: CreateCourseType :one
-- Creates a new course type
INSERT INTO course_types (name, rate) VALUES ($1, $2) RETURNING *;

-- name: GetCourseType :one
-- Retrieves a course type by ID
SELECT * FROM course_types WHERE id = $1;

-- name: ListCourseTypes :many
-- Lists all course types
SELECT * FROM course_types;

-- name: UpdateCourseType :one
-- Updates a course type's information
UPDATE course_types SET name = $2, rate = $3 WHERE id = $1 RETURNING *;

-- name: DeleteCourseType :exec
-- Deletes a course type
DELETE FROM course_types WHERE id = $1;
