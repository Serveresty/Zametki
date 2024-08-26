package database

import (
	"Zametki/models"
	"context"
	"time"
)

func CreateNoteDB(note *models.Note, id string) error {
	note.CreatedAt = time.Now()
	_, err := db.Exec(context.Background(), `INSERT INTO "notes" (user_id, title, content, created_at) VALUES($1,$2,$3,$4)`,
		id, note.Title, note.Content, note.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
