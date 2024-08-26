package database

import (
	"Zametki/models"
	"context"
)

func GetNotesDB(id string) ([]models.Note, error) {
	rows, err := db.Query(context.Background(), `SELECT note_id, title, content, created_at FROM notes WHERE user_id = $1`, id)
	if err != nil {
		return []models.Note{}, err
	}

	var notes []models.Note
	for rows.Next() {
		var note models.Note

		err = rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt)
		if err != nil {
			return []models.Note{}, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
