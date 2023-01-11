package handlers

import (
	"encoding/json"
	"log"
	sharedErrors "shared/errors"
	"time"

	"authentication/communications"
	"authentication/global"
	"authentication/middlewares"
	proto "authentication/proto/generated"
	"authentication/types"
)

func StartRegistrationHandler(
	startRegistrationRequest *proto.StartRegistrationRequest) (*proto.StartRegistrationResponse, error) {

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

	error= global.GlobalVariables.RedisClient.Set(startRegistrationRequest.Email, temporaryUserDetails, 600 * time.Second).Err( )
	if error != nil {
		log.Println(error.Error( ))

		return &proto.StartRegistrationResponse{ Error: &sharedErrors.ServerError }, nil }

	//! sending request to otp service to send OTP to the email for verification
	error= communications.SendEmailVerificationOTP(startRegistrationRequest.Email)
	if error != nil {
		return &proto.StartRegistrationResponse{ Error: &sharedErrors.ServerError }, nil }

	return nil, nil
}