package database

import (
	"Zametki/configs"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func InitDB() error {
	username := configs.GetEnv("DB_USERNAME")
	password := configs.GetEnv("DB_PASSWORD")
	host := configs.GetEnv("DB_HOST")
	port := configs.GetEnv("DB_PORT")
	dbName := configs.GetEnv("DB_NAME")

	dbUrl := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + dbName

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return fmt.Errorf("error while connecting to database: %v", err)
	}

	db = conn

	err = createBaseTables()
	if err != nil {
		return err
	}

	return nil
}

func createBaseTables() error {
	_, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS "users" (user_id serial PRIMARY KEY, first_name VARCHAR(50) NOT NULL, last_name VARCHAR(50) NOT NULL, email VARCHAR(255) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL);
	CREATE TABLE IF NOT EXISTS "roles" (role_id serial PRIMARY KEY, role_name VARCHAR(20) UNIQUE NOT NULL);
	CREATE TABLE IF NOT EXISTS "users_to_roles" (user_id int references users (user_id) on delete cascade, role_id int references roles (role_id) on delete cascade);
	CREATE TABLE IF NOT EXISTS "notes" (note_id serial PRIMARY KEY, user_id int references users (user_id) on delete cascade, title VARCHAR(50) NOT NULL, content VARCHAR(255) NOT NULL, created_at TIMESTAMP NOT NULL)`)
	if err != nil {
		return fmt.Errorf("error while creating base tables: %v", err)
	}

	_, _ = db.Exec(context.Background(), `
	INSERT INTO "roles" ("role_name") VALUES ('user'), ('admin');
	`)

	return nil
}
