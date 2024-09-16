-- name: CreateCourseOffering :one
-- Creates a new course offering
INSERT INTO course_offerings (name, starting_date, ending_date, price, sport_center_id, course_type_id, instructor_id)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetCourseOffering :one
-- Retrieves a course offering by ID
SELECT * FROM course_offerings WHERE id = $1;

-- name: ListCourseOfferingsBySportCenter :many
-- Lists all course offerings for a sport center
SELECT * FROM course_offerings WHERE sport_center_id = $1;

-- name: UpdateCourseOffering :one
-- Updates a course offering's information
UPDATE course_offerings
SET name = $2, starting_date = $3, ending_date = $4, price = $5, course_type_id = $6, instructor_id = $7
WHERE id = $1 RETURNING *;

-- name: DeleteCourseOffering :exec
-- Deletes a course offering
DELETE FROM course_offerings WHERE id = $1;
