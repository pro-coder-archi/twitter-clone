package server

import (
	"context"

	"authentication/handlers"
	"authentication/proto"
)

//* server implementation for gRPC service Authentication
type AuthenticationServer struct {
	proto.UnimplementedAuthenticationServer
}

func(server *AuthenticationServer) StartRegistration(
	ctx context.Context, startRegistrationRequest *proto.StartRegistrationRequest) (*proto.StartRegistrationResponse, error) {

	return handlers.StartRegistrationHandler(startRegistrationRequest)
}