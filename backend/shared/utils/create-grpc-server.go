package utils

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var gRPCServerPort= flag.Int("port", 4000, "Port where gRPC server will listen")

func CreateGRPCServer(setupGrpcServer func(*grpc.Server)) {

	//* creating a tcp listener which will listen at port 4000
	portListener, error := net.Listen("tcp", fmt.Sprintf("localhost:%d", *gRPCServerPort))

	if error != nil {
		log.Fatalf("❌ error listening on port %d", *gRPCServerPort)

		log.Fatalln(error.Error( )) }

	//* creating the grpc server

	gRPCServer := grpc.NewServer( )
	reflection.Register(gRPCServer)

	setupGrpcServer(gRPCServer)

	//* starting the grpc server

	log.Println("🚀 starting gRPC server for authentication microservice !")

	gRPCServer.Serve(portListener)

	defer gRPCServer.Stop( )
}