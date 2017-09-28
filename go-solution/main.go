package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/fnugk/interview-assignment/go-solution/db"
	"github.com/pkg/errors"
)

var (
	dbPath    = "./test.db"
	outFolder = "./out"
)

func main() {
	if err := os.MkdirAll(outFolder, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.NewDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	userDB := db.NewUserDB(dbConn)

	ctx := context.Background()
	err = dbConn.Tx(ctx, func(ctx context.Context, tx *sql.Tx) error {
		users, err := userDB.GetAll(ctx, tx)
		if err != nil {
			return err
		}

		for i, u := range users {
			buf, err := json.Marshal(u)
			if err != nil {
				return errors.Wrapf(err, "could marshal %v", u)
			}
			fileName := fmt.Sprintf("user-%d.json", i)
			err = writeAndVerify(buf, path.Join(outFolder, fileName))
			if err != nil {
				return err
			}
			log.Println("wrote", fileName)
			err = userDB.DeleteByID(ctx, tx, u.ID)
			if err != nil {
				return err
			}
			log.Println("deleted", u)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func writeAndVerify(content []byte, fPath string) error {
	f, err := os.Create(fPath)
	if err != nil {
		return errors.Wrap(err, "could not create file")
	}
	defer f.Close()

	n, err := f.Write(content)
	if err != nil {
		return errors.Wrap(err, "could not write to file")
	}
	if n != len(content) {
		return errors.Errorf("only wrote %d/%d bytes", n, len(content))
	}
	return nil
}
