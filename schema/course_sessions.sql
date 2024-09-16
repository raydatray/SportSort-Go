-- name: CreateCourseSession :one
-- Creates a new course session
INSERT INTO course_sessions (date, course_offering_id, room_id, instructor_id, start_time, end_time)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetCourseSession :one
-- Retrieves a course session by ID
SELECT * FROM course_sessions WHERE id = $1;

-- name: ListCourseSessionsByCourseOffering :many
-- Lists all course sessions for a course offering
SELECT * FROM course_sessions WHERE course_offering_id = $1;

-- name: UpdateCourseSession :one
-- Updates a course session's information
UPDATE course_sessions SET date = $2, room_id = $3, instructor_id = $4, start_time = $5, end_time = $6
WHERE id = $1 RETURNING *;

-- name: DeleteCourseSession :exec
-- Deletes a course session
DELETE FROM course_sessions WHERE id = $1;
