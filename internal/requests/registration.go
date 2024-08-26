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
	_, err := jwts.ParseToken(token)
	if err == nil {
		RespondWithError(w, http.StatusBadRequest, customerrors.ErrAlreadyAuthorized.Error())
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = database.IsUserRegistered(user.Email)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = database.RegistationDB(&user)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
