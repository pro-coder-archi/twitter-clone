package main

import (
	"os"
	sharedUtils "shared/utils"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"

	"authentication/configuration"
	"authentication/globals"
	proto "authentication/proto/generated"
	"authentication/repository"
	"authentication/server"
)

func main( ) {
	var (
		methodName= "main"

		cleanup func( )
	)

	//! validating app configuration
	var appConfiguration configuration.AppConfiguration
	error := envconfig.Process("", &appConfiguration)

	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Method: methodName,
			Message: error,
		})

		os.Exit(1)
	}

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