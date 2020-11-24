package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Token ...
type Token struct {
	UserID string
	jwt.StandardClaims
}