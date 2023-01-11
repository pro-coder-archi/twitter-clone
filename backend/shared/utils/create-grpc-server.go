package utils

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var gRPCServerPort= flag.Int("port", 4000, "Port where gRPC server will listen")

func CreateGRPCServer(setupGrpcServer func(*grpc.Server)) {
	const methodName= "CreateGRPCServer"

	//* creating a tcp listener which will listen at port 4000
	portListener, error := net.Listen("tcp", fmt.Sprintf("localhost:%d", *gRPCServerPort))

	if error != nil {
		Log(LogDetails{

			Method: methodName,
			Message: fmt.Sprintf("❌ error listening on port %d\n%s", *gRPCServerPort, error.Error( )),
		})}

	//* creating the grpc server

	gRPCServer := grpc.NewServer( )
	reflection.Register(gRPCServer)

	setupGrpcServer(gRPCServer)

	//* starting the grpc server

	Log(LogDetails{

		Type: DEFAULT_LOG,
		Method: methodName,
		Message: "🚀 starting gRPC server for authentication microservice !",
	})

	gRPCServer.Serve(portListener)

	defer gRPCServer.Stop( )
}