package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"shared/errors"
)

type JwtPayload struct {
	jwt.RegisteredClaims

	Email string `json:"email"`
}

const jwtSigningSecert= "secret"

func CreateJwt(email string) (string, *string) {
	const methodName= "CreateJwt"

	jwtPayload := JwtPayload{
		Email: email,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now( ).Add(time.Hour * 24)),
		},
	}

	token, error := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtPayload).
		SignedString([]byte(jwtSigningSecert))

	if error != nil {
		Log(LogDetails{

			Method: methodName,
			Message: error,
		})

		return token, &errors.ServerError }

	return token, nil
}

func VerifyJwt(token string) (bool, *string) {
	const methodName= "VerifyJwt"

	var (
		jwtPayload JwtPayload
		verifyJwtError string
	)

	_, error := jwt.ParseWithClaims(
		token, &jwtPayload,
			func(t *jwt.Token) (interface{ }, error) {
				return []byte(jwtSigningSecert), nil
			},
	)

	if error != nil {

		if error == jwt.ErrSignatureInvalid {
			verifyJwtError= "jwt expired or not found" } else {

		Log(LogDetails{

			Method: methodName,
			Message: error,
		})

		verifyJwtError= errors.ServerError }

		return false, &verifyJwtError
	}

	return true, nil
}