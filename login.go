package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"log"

	"github.com/jmoiron/sqlx"
)

//Login handles the user login
func Login(ctx context.Context, db *sqlx.DB, user User) (ok bool, err error) {
	// Begin db transaction
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Query for salt
	var salt string
	err = tx.QueryRowxContext(
		ctx,
		`
			SELECT salt FROM users WHERE username = $1;
		`,
		user.Username,
	).Scan(&salt)
	if err == sql.ErrNoRows {
		return
	} else if err != nil {
		return
	}

	// Get the hash from the password and salt
	s, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return
	}
	hash := HashWithSalt([]byte(user.Password), s)
	if err != nil {
		return
	}

	// Get the user from the username and hash
	var id int
	err = tx.QueryRowxContext(
		ctx,
		`
			SELECT id FROM users WHERE username = $1 AND hash = $2
		`,
		user.Username,
		base64.StdEncoding.EncodeToString(hash),
	).Scan(&id)
	if err == sql.ErrNoRows {
		log.Println("Invalid login")
	} else if err != nil {
		log.Fatal(err)
	} else {
		return true, nil
	}
	return
}
