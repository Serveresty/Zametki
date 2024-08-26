package jwts

import (
	"Zametki/configs"
	"Zametki/models"
	cerr "Zametki/utils/custom-errors"
	customerrors "Zametki/utils/custom-errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token string) (*models.Claims, error) {
	if token == "" {
		return nil, cerr.ErrNoTokenFound
	}

	tokenString := strings.Split(token, " ")[1]

	secretKey := []byte(configs.GetEnv("SECRET_KEY"))
	tkn, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !tkn.Valid {
		return nil, cerr.ErrInvalidToken
	}

	claims, ok := tkn.Claims.(*models.Claims)
	if !ok {
		return nil, customerrors.ErrClaims
	}

	return claims, nil
}
