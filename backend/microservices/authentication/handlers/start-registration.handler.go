package handlers

import (
	"encoding/json"

	"authentication/communications"
	"authentication/global"
	"authentication/middlewares"
	"authentication/proto"
)

func StartRegistrationHandler(
	startRegistrationRequest *proto.StartRegistrationRequest) (*proto.StartRegistrationResponse, error) {

	//! request passing through middleware

	middlewareFiltrationError := middlewares.StartRegistrationMiddleware(startRegistrationRequest)

	if middlewareFiltrationError != nil {
		return &proto.StartRegistrationResponse{ Error: middlewareFiltrationError }, nil }

	//! saving the details temporarily in redis

	temporaryUserDetails, error := json.Marshal(startRegistrationRequest)
	if error != nil {
		return &proto.StartRegistrationResponse{ Error: &global.ServerError }, nil }

	error= global.GlobalVariables.RedisClient.Set(startRegistrationRequest.Email, temporaryUserDetails, 600).Err( )
	if error != nil {
		return &proto.StartRegistrationResponse{ Error: &global.ServerError }, nil }

	//! sending request to otp service to send OTP to the email for verification
	error= communications.SendEmailVerificationOTP(startRegistrationRequest.Email)
	if error != nil {
		return &proto.StartRegistrationResponse{ Error: &global.ServerError }, nil }

	return &proto.StartRegistrationResponse{ }, nil
}