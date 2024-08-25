package jwts

import (
	"Zametki/configs"
	"Zametki/models"
	cerr "Zametki/utils/custom-errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string) error {
	if token == "" {
		return cerr.ErrNoTokenFound
	}

	tokenString := strings.Split(token, " ")[1]

	secretKey := []byte(configs.GetEnv("SECRET_KEY"))
	tkn, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !tkn.Valid {
		return cerr.ErrInvalidToken
	}

	return nil
}
