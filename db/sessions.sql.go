// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sessions.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const clearSessions = `-- name: ClearSessions :exec
DELETE FROM sessions WHERE expires_at < CURRENT_TIMESTAMP
`

// Clears expired sessions
func (q *Queries) ClearSessions(ctx context.Context) error {
	_, err := q.db.Exec(ctx, clearSessions)
	return err
}

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (user_id, session_token, created_at, expires_at)
VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '24 hours')
RETURNING id, session_token, expires_at
`

type CreateSessionParams struct {
	UserID       int64
	SessionToken string
}

type CreateSessionRow struct {
	ID           int64
	SessionToken string
	ExpiresAt    pgtype.Timestamptz
}

// Creates a new session for a given use
func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (CreateSessionRow, error) {
	row := q.db.QueryRow(ctx, createSession, arg.UserID, arg.SessionToken)
	var i CreateSessionRow
	err := row.Scan(&i.ID, &i.SessionToken, &i.ExpiresAt)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions WHERE session_token = $1
`

// Delete a session when logging out
func (q *Queries) DeleteSession(ctx context.Context, sessionToken string) error {
	_, err := q.db.Exec(ctx, deleteSession, sessionToken)
	return err
}

const refreshSession = `-- name: RefreshSession :one
UPDATE sessions
SET expires_at = CURRENT_TIMESTAMP + INTERVAL '24 hours'
WHERE session_token = $1
RETURNING id, session_token, expires_at
`

type RefreshSessionRow struct {
	ID           int64
	SessionToken string
	ExpiresAt    pgtype.Timestamptz
}

// Refreshes a user's session
func (q *Queries) RefreshSession(ctx context.Context, sessionToken string) (RefreshSessionRow, error) {
	row := q.db.QueryRow(ctx, refreshSession, sessionToken)
	var i RefreshSessionRow
	err := row.Scan(&i.ID, &i.SessionToken, &i.ExpiresAt)
	return i, err
}

const validateSession = `-- name: ValidateSession :exec
SELECT s.id, s.user_id, u.type AS user_type, s.expires_at
FROM sessions s
JOIN users u ON s.user_id = u.id
WHERE s.session_token = $1 AND s.expires_at > CURRENT_TIMESTAMP
`

// Validates a user's session
func (q *Queries) ValidateSession(ctx context.Context, sessionToken string) error {
	_, err := q.db.Exec(ctx, validateSession, sessionToken)
	return err
}
