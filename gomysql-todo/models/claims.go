package models

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	Username string
	UserID   uint
	jwt.RegisteredClaims
}
