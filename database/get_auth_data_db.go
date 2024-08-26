package database

import (
	"Zametki/models"
	"Zametki/utils/password"
	"context"
	"fmt"
)

func GetAuthData(user *models.User) ([]string, error) {
	paswrd := user.Password

	row := db.QueryRow(context.Background(), `SELECT user_id, first_name, last_name, password FROM users WHERE email = $1`, user.Email)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return []string{}, err
	}

	if !password.CheckPasswordHash(paswrd, user.Password) {
		return []string{}, fmt.Errorf("bad credentials")
	}

	rows, err := db.Query(context.Background(), `SELECT role_name FROM roles JOIN users_to_roles ON roles.role_id=users_to_roles.role_id WHERE users_to_roles.user_id=$1`, user.ID)
	if err != nil {
		return []string{}, err
	}

	var roles []string
	for rows.Next() {
		var role string

		err = rows.Scan(&role)
		if err != nil {
			return []string{}, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
