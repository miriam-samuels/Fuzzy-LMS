package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type SignUpCred struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type SignInCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	UserId string `json:"_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
