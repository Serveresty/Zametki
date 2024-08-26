package requests

import (
	"Zametki/database"
	"Zametki/models"
	customerrors "Zametki/utils/custom-errors"
	"Zametki/utils/jwts"
	"encoding/json"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	_, err := jwts.ParseToken(token)
	if err == nil {
		http.Error(w, customerrors.ErrAlreadyAuthorized.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	roles, err := database.GetAuthData(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idString := strconv.Itoa(user.ID)
	token, err = jwts.CreateToken(idString, roles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenJSON := map[string]string{
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", token)
	json.NewEncoder(w).Encode(tokenJSON)
	w.WriteHeader(http.StatusOK)
}
