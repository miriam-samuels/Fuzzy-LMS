package types

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SignUpCred struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type SignInCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	UserId uuid.UUID `json:"_id"`
	jwt.RegisteredClaims
}
