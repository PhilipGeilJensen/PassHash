package main

import (
	"context"
	"encoding/base64"

	"github.com/jmoiron/sqlx"
)

//Create the user
func CreateUser(ctx context.Context, user User, db *sqlx.DB) (ok bool, err error) {
	salt := GenerateSalt()
	hash := HashWithSalt([]byte(user.Password), salt)

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	_, err = tx.ExecContext(
		ctx,
		`
			INSERT INTO users (username, hash, salt)
			VALUES ($1, $2, $3)
		`,
		user.Username,
		base64.StdEncoding.EncodeToString(hash),
		base64.StdEncoding.EncodeToString(salt),
	)
	if err != nil {
		return
	}

	tx.Commit()
	return true, err
}
