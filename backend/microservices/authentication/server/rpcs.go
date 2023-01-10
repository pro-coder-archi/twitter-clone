package server

import (
	"context"

	"authentication/handlers"
	proto "authentication/proto/generated"
)

func(server *AuthenticationServer) StartRegistration(
	ctx context.Context, startRegistrationRequest *proto.StartRegistrationRequest) (*proto.StartRegistrationResponse, error) {

	return handlers.StartRegistrationHandler(startRegistrationRequest)
}

func(server *AuthenticationServer) SetEmailVerified(
	ctx context.Context, setEmailVerifiedRequest *proto.SetEmailVerifiedRequest) (*proto.SetEmailVerifiedResponse, error) {

	return handlers.SetEmailVerifiedHandler(setEmailVerifiedRequest)
}