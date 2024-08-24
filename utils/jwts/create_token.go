package jwts

import (
	"Zametki/configs"
	"Zametki/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id string, roles []string) (string, error) {
	claims := &models.Claims{
		Roles: roles,
		StandardClaims: jwt.StandardClaims{
			Subject:   id,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.GetEnv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return "Bearer " + tokenString, nil
}
