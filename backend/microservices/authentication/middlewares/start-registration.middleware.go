package middlewares

import (
	"context"
	"database/sql"
	"net/mail"
	sharedErrors "shared/errors"

	"authentication/global"
	proto "authentication/proto/generated"
)

func StartRegistrationMiddleware(startRegistrationRequest *proto.StartRegistrationRequest) *string {

	//! basic validations

	_, error := mail.ParseAddress(startRegistrationRequest.Email)

	if error != nil {
		return &InvalidEmailError }

	if len(startRegistrationRequest.Name) < 3 || len(startRegistrationRequest.Name) > 50 {
		return &InvalidNameError }

	//! check if email is not pre-registered using the authentication database

	queryResult, error := global.GlobalVariables.Repository.FindRegisteredEmail(
		context.Background( ), startRegistrationRequest.Email,
	)

	if queryResult.Email == startRegistrationRequest.Email {
		return &DuplicateEmailError }

	if error == sql.ErrNoRows { return nil }

	return &sharedErrors.ServerError
}

var (
	InvalidEmailError= "email is invalid"
	InvalidNameError= "name should be 3 to 50 characters long"
	DuplicateEmailError= "email is pre-registered"
)