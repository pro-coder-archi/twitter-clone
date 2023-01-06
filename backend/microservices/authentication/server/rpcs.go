package server

import (
	"context"

	"authentication/handlers"
	"authentication/proto"
)

func(server *AuthenticationServer) StartRegistration(
	ctx context.Context, startRegistrationRequest *proto.StartRegistrationRequest) (*proto.StartRegistrationResponse, error) {

	return handlers.StartRegistrationHandler(startRegistrationRequest)
}