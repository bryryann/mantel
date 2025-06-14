package data

import (
	"context"
	"database/sql"
	"time"
)

type Follows struct {
	ID         int64     `json:"id"`
	FollowerID int64     `json:"follower_id"`
	FolloweeID int64     `json:"followee_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type FollowsModel struct {
	DB *sql.DB
}

// Insert adds a new follow record to the database.
func (m FollowsModel) Insert(followerID, followeeID int64) error {
	query := `
		INSERT INTO follows (follower_id, followee_id)
		VALUES ($1, $2)
		ON CONFLICT (follower_id, followee_id) DO NOTHING`

	args := []any{followerID, followeeID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...)
	if err != nil {
		// TODO: Add clearer error messages
		return err.Err()
	}

	return nil
}

// Delete removes a follow record from the follow table.
func (m FollowsModel) Delete(followerID, followeeID int64) error {
	query := `
		DELETE FROM follows
		WHERE follower_id = $1 AND followee_id = $2`

	args := []any{followerID, followeeID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...)
	if err != nil {
		return err.Err()
	}

	return nil
}
