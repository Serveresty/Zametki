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
		RespondWithError(w, http.StatusBadRequest, customerrors.ErrAlreadyAuthorized.Error())
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	roles, err := database.GetAuthData(&user)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	idString := strconv.Itoa(user.ID)
	token, err = jwts.CreateToken(idString, roles)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
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
