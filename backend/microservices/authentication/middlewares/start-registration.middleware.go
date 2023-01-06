package middlewares

import (
	"context"
	"database/sql"
	"net/mail"

	"authentication/global"
	"authentication/proto"
)

func StartRegistrationMiddleware(startRegistrationRequest *proto.StartRegistrationRequest) *string {

	//! input validation

	// email address validation
	_, error := mail.ParseAddress(startRegistrationRequest.Email)

	if error != nil {
		return &invalidEmailError }

	// name validation
	if len(startRegistrationRequest.Name) < 3 || len(startRegistrationRequest.Name) > 50 {
		return &invalidNameError }

	//! check if email is not pre-registered using the authentication database

	queryResult, error := global.GlobalVariables.Repository.FindRegisteredEmail(
		context.Background( ), startRegistrationRequest.Email,
	)

	// handling duplicate email error
	if queryResult.Email == startRegistrationRequest.Email {
		return &duplicateEmailError }

	// handling server or database error
	if error == sql.ErrNoRows { return nil }

	return &global.ServerError
}

var (
	invalidEmailError= "email is invalid"
	invalidNameError= "name should be 3 to 50 characters long"
	duplicateEmailError= "email is pre-registered"
)