package handlers

import (
	"encoding/json"
	sharedErrors "shared/errors"
	sharedUtils "shared/utils"
	"time"

	"authentication/communications"
	"authentication/globals"
	"authentication/middlewares"
	proto "authentication/proto/generated"
	"authentication/types"
)

func StartRegistrationHandler(
	startRegistrationRequest *proto.StartRegistrationRequest) (*proto.StartRegistrationResponse, error) {

	const methodName= "StartRegistrationHandler"

	//! request passing through middleware

	middlewareFiltrationError := middlewares.StartRegistrationMiddleware(startRegistrationRequest)

	if middlewareFiltrationError != nil {
		return &proto.StartRegistrationResponse{ Error: middlewareFiltrationError }, nil }

	//! saving the details temporarily in redis

	temporaryUserDetails, error := json.Marshal(
		types.TemporaryUserDetailsRedisRecord {
			IsVerified: false,

			Email: startRegistrationRequest.Email,
			Name: startRegistrationRequest.Name,
		},
	)

	if error != nil {
		return &proto.StartRegistrationResponse{ Error: &sharedErrors.ServerError }, nil }

	error= globals.Variables.RedisClient.Set(startRegistrationRequest.Email, temporaryUserDetails, 600 * time.Second).Err( )
	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Message: error,
			Method: methodName,
		})

		return &proto.StartRegistrationResponse{ Error: &sharedErrors.ServerError }, nil }

	//! sending request to otp service to send OTP to the email for verification
	error= communications.SendEmailVerificationOTP(startRegistrationRequest.Email)
	if error != nil {
		return &proto.StartRegistrationResponse{ Error: &sharedErrors.ServerError }, nil }

	return nil, nil
}