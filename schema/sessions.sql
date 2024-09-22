-- name: CreateSession :one
-- Creates a new session for a given use
INSERT INTO sessions (user_id, session_token, created_at, expires_at)
VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '24 hours')
RETURNING id, session_token, expires_at;

-- name: ValidateSession :exec
-- Validates a user's session
SELECT s.id, s.user_id, u.type AS user_type, s.expires_at
FROM sessions s
JOIN users u ON s.user_id = u.id
WHERE s.session_token = $1 AND s.expires_at > CURRENT_TIMESTAMP;

-- name: RefreshSession :one
-- Refreshes a user's session
UPDATE sessions
SET expires_at = CURRENT_TIMESTAMP + INTERVAL '24 hours'
WHERE session_token = $1
RETURNING id, session_token, expires_at;

-- name: DeleteSession :exec
-- Delete a session when logging out
DELETE FROM sessions WHERE session_token = $1;

-- name: ClearSessions :exec
-- Clears expired sessions
DELETE FROM sessions WHERE expires_at < CURRENT_TIMESTAMP;
