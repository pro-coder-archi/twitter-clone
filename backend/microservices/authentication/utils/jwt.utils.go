package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"authentication/global"
)

type JwtPayload struct {
	jwt.StandardClaims

	Email string `json:"email"`
}

const jwtSigningSecert= "secret"

func CreateJwt(email string) (string, *string) {

	jwtPayload := JwtPayload{
		Email: email,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now( ).Add(time.Hour * 24).Unix( ),
		},
	}

	token, error := jwt.NewWithClaims(jwt.SigningMethodES256, &jwtPayload).
		SignedString([]byte(jwtSigningSecert))

	if error != nil {
		return token, &global.ServerError}

	return token, nil
}

func VerifyJwt(token string) (bool, *string) {

	var (
		jwtPayload JwtPayload
		verifyJwtError string
	)

	_, error := jwt.ParseWithClaims(
		token, &jwtPayload,
			func(t *jwt.Token) (interface{ }, error) {
				return jwtSigningSecert, nil
			},
	)

	if error != nil {

		if error == jwt.ErrSignatureInvalid {
			verifyJwtError= "jwt expired or not found"} else {

		verifyJwtError= global.ServerError}

		return false, &verifyJwtError
	}

	return true, nil
}