package requests

import (
	"Zametki/database"
	"Zametki/models"
	customerrors "Zametki/utils/custom-errors"
	"Zametki/utils/jwts"
	"encoding/json"
	"net/http"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	err := jwts.ParseToken(token)
	if err == nil {
		http.Error(w, customerrors.ErrAlreadyAuthorized.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.IsUserRegistered(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.RegistationDB(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
