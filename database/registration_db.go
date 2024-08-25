package database

import (
	"Zametki/models"
	"Zametki/utils/password"
	"context"
	"log"
)

func RegistationDB(user *models.User) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if rerr := tx.Rollback(context.Background()); rerr != nil {
				log.Printf("rollback error: %v", rerr)
			}
		}
	}()

	user.Password, err = password.HashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `INSERT INTO "users" (first_name, last_name, email, password) VALUES($1,$2,$3,$4)`,
		user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	row := tx.QueryRow(context.Background(), `SELECT user_id FROM "users" WHERE email = $1`, user.Email)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `INSERT INTO "users_to_roles" (user_id, role_id) VALUES($1,$2)`,
		id, 1)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
