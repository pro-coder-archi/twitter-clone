package server

import (
	proto "authentication/proto/generated"
)

//* server implementation for gRPC service Authentication
type AuthenticationServer struct {
	proto.UnimplementedAuthenticationServer
}