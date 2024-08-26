package requests

import (
	"Zametki/database"
	"Zametki/utils/jwts"
	"encoding/json"
	"net/http"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	claims, err := jwts.ParseToken(token)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	notes, err := database.GetNotesDB(claims.Subject)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
	w.WriteHeader(http.StatusOK)
}
