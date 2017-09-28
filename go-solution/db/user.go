package db

import (
	"context"
	"database/sql"

	"github.com/fnugk/interview-assignment/go-solution/model"
	"github.com/pkg/errors"
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

// DeleteByID deletes the user by the given id
func (db *UserDB) DeleteByID(ctx context.Context, tx *sql.Tx, id string) error {
	q := `
		DELETE FROM useres
		WHERE id=?
	`

	res, err := tx.ExecContext(ctx, q, id)
	if err != nil {
		return errors.Wrapf(err, "could not delete users with id %#v", id)
	}

	// though unique id is not enforced by the db, this assumes unique id
	n, err := res.RowsAffected()
	if err != nil {
		return errors.Wrapf(err, "could not get RowsAffected when deleteing by id %#v", id)
	}

	if n != 1 {
		return errors.Errorf("expected to delete exactly 1 row instead deleted %d rows", n)
	}
	return nil
}
