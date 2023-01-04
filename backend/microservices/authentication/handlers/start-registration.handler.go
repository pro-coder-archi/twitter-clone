package handlers

import (
	"context"
	"encoding/json"

	"authentication/global"
	"authentication/proto"
)

func StartRegistrationHandler(
	startRegistrationRequest *proto.StartRegistrationRequest) (*proto.StartRegistrationResponse, error) {

	//* check if email is not pre-registered using the authentication database

	queryResult, error := global.GlobalVariables.Repository.FindRegisteredEmail(
		context.Background( ), startRegistrationRequest.Email,
	)

	// handling duplicate email error
	if(queryResult.Email == startRegistrationRequest.Email) {

		duplicateEmailError := "email is pre-registered"

		return &proto.StartRegistrationResponse{
			Error: &duplicateEmailError,
		}, nil
	}

	// handling server or database error
	if(error != nil) { return nil, global.ServerError }

	//* saving the details in redis

	temporaryUserDetails, error := json.Marshal(startRegistrationRequest)
	if(error != nil) { return nil, global.ServerError }

	error= global.GlobalVariables.RedisClient.Set(startRegistrationRequest.Email, temporaryUserDetails, 600).Err( )
	if(error != nil) { return nil, global.ServerError }

	// TODO: send request to otp service to send OTP to the email

	return &proto.StartRegistrationResponse{ }, nil
}