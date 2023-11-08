package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/miriam-samuels/loan-management-backend/internal/constants"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

func SignJWT(user string, role string) (string, error) {
	// generate new token with signing method and claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, types.JWTClaims{
		UserId: user,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			//  issued at the utc equivalent of the current time
			IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
			//  token to expire in 6hrs from the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * 6)),
			Issuer:    constants.Env.AppName,
		},
	})

	// return token with an error message if it fails
	return claims.SignedString([]byte(constants.Env.JWTSecret))
}

func VerifyJWT(token string) (*types.JWTClaims, bool) {
	// Parse the token string and store the result in claims
	tkn, err := jwt.ParseWithClaims(token, &types.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(constants.Env.JWTSecret), nil
	})

	// if no error and token is valid return thr claims and a true for validity
	if err == nil && tkn.Valid {
		return tkn.Claims.(*types.JWTClaims), true
	}

	return nil, false
}
