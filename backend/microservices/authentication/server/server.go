package server

import (
	"authentication/proto"
)

//* server implementation for gRPC service Authentication
type AuthenticationServer struct {
	proto.UnimplementedAuthenticationServer
}