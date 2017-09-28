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
	db *DB
}

// NewUserDB creates a new UserDB
func NewUserDB(db *DB) *UserDB {
	return &UserDB{db}
}

// DeleteByID deletes the user by the given id
func (db *UserDB) DeleteByID(ctx context.Context, tx *sql.Tx, id string) error {
	q := `
		DELETE FROM users
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

// GetAll returns a list of all users
func (db *UserDB) GetAll(ctx context.Context, tx *sql.Tx) ([]*model.User, error) {
	q := `
		SELECT id, firstName, lastName, email
		FROM users
	`

	rows, err := tx.QueryContext(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err, "could not GetAll users")
	}
	defer rows.Close()

	users := []*model.User{}
	for rows.Next() {
		var usr model.User
		err := rows.Scan(&usr.ID, &usr.FirstName, &usr.LastName, &usr.Email)
		if err != nil {
			return nil, errors.Wrap(err, "could not scan user row")
		}
		users = append(users, &usr)
	}

	return users, rows.Err()
}
