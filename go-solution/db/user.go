package db

import (
	"context"
	"database/sql"

	"github.com/fnugk/interview-assignment/go-solution/model"
)

// IUserDB exposes all functionality of a UserDB
type IUserDB interface {
	DeleteByID(ctx context.Context, tx *sql.Tx, id string) error
	GetAll(ctx context.Context, tx *sql.Tx) ([]*model.User, error)
}

// UserDB abstracts the users table
type UserDB struct {
	*DB
}

// NewUserDB creates a new UserDB
func NewUserDB(db *DB) *UserDB {
	return &UserDB{db}
}
