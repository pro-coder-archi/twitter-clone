package main

import (
	sharedUtils "shared/utils"

	"google.golang.org/grpc"

	"authentication/globals"
	proto "authentication/proto/generated"
	"authentication/repository"
	"authentication/server"
)

func main( ) {
	var cleanup func( )

	//! connecting to cockroachDB
	dbConnection := sharedUtils.CreateCockroachDBConnection( )
	defer dbConnection.Close( )

	globals.Variables.Repository = repository.New(dbConnection)

	//! connecting to redis

	globals.Variables.RedisClient, cleanup= sharedUtils.CreateRedisClient(false)
	defer cleanup( )

	//! starting the gRPC server
	sharedUtils.CreateGRPCServer(
		func(gRPCServer *grpc.Server) {
			proto.RegisterAuthenticationServer(gRPCServer, &server.AuthenticationServer{ })
		},
	)
}