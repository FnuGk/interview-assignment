package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Needed so sql.Open knows what a sqlite3 db is
	"github.com/pkg/errors"
)

// DB is the main database handle
type DB struct {
	*sql.DB
}

// NewDB connects to a sqlite3 database and returns a handle
func NewDB(dbPath string) (*DB, error) {
	connectionString := fmt.Sprintf("%s", dbPath) // sqlite3 has lots of connection options that can be set here
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, errors.Wrapf(err, "could not connect to %#v", connectionString)
	}
	// Actual connection is only guarenteed after ping
	if err := db.Ping(); err != nil {
		return nil, errors.Wrapf(err, "could not ping db at %#v", connectionString)
	}
	return &DB{db}, nil
}

// Tx runs the function f in a transaction that is commited if err is nil, else it will roll back
func (db *DB) Tx(ctx context.Context, f func(ctx context.Context, tx *sql.Tx) error) error {
	return transact(ctx, db.DB, f)
}

// transact is a helper function for managing transactions
func transact(ctx context.Context, db *sql.DB, f func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = f(ctx, tx)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}
