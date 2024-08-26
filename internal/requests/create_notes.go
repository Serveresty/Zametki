package requests

import (
	"Zametki/database"
	"Zametki/models"
	"Zametki/utils/jwts"
	"encoding/json"
	"net/http"
)

func CreateNotes(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	claims, err := jwts.ParseToken(token)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = database.CreateNoteDB(&note, claims.Subject)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
