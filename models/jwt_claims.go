package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Roles []string
	jwt.StandardClaims
}
