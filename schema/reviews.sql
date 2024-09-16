-- name: CreateReview :one
-- Creates a new review
INSERT INTO reviews (user_id, rating, comment, review_date, instructor_id, course_type_id)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetReview :one
-- Retrieves a review by ID
SELECT * FROM reviews WHERE id = $1;

-- name: ListUserReviews :many
-- Lists all reviews by a user
SELECT * FROM reviews WHERE user_id = $1;


-- name: ListInstructorReviews :many
-- Lists all reviews for an instructor
SELECT * FROM reviews WHERE instructor_id = $1;

-- name: ListCourseTypeReviews :many
-- Lists all reviews for a course type
SELECT * FROM reviews WHERE course_type_id = $1;

-- name: UpdateReview :one
-- Updates a review's information
UPDATE reviews SET rating = $2, comment = $3 WHERE id = $1 RETURNING *;

-- name: DeleteReview :exec
-- Deletes a review
DELETE FROM reviews WHERE id = $1;
