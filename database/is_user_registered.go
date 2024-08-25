package database

import (
	customerrors "Zametki/utils/custom-errors"
	"context"

	"github.com/jackc/pgx/v4"
)

func IsUserRegistered(email string) error {
	row := db.QueryRow(context.Background(), `SELECT user_id FROM users WHERE email = $1`, email)
	var a int
	err := row.Scan(&a)
	switch err {
	case nil:
		return customerrors.ErrAlreadyRegistered
	case pgx.ErrNoRows:
		return nil
	default:
		return err
	}
}
