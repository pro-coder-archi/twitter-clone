package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"authentication/proto"
	"authentication/server"
)

var gRPCServerPort= flag.Int("port", 4000, "Port where gRPC server will listen")

func main( ) {

	portListener, error := net.Listen("tcp", fmt.Sprintf("localhost:%d", *gRPCServerPort))

	if error != nil {
		log.Fatalf("❌ error listening on port %d", *gRPCServerPort)

		log.Fatalln(error.Error( )) }

	gRPCServer := grpc.NewServer( )
	proto.RegisterAuthenticationServer(gRPCServer, &server.AuthenticationServer{ })

	log.Println("🚀 starting gRPC server for authentication microservice !")

	gRPCServer.Serve(portListener)
}