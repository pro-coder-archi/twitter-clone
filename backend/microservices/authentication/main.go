package main

import (
	sharedUtils "shared/utils"

	"authentication/global"
	"authentication/repository"
)

func main( ) {
	var cleanup func( )

	//! connecting to cockroachDB
	dbConnection := sharedUtils.CreateCockroachDBConnection( )
	defer dbConnection.Close( )

	global.GlobalVariables.Repository = repository.New(dbConnection)

	//! connecting to redis

	global.GlobalVariables.RedisClient, cleanup= sharedUtils.CreateRedisClient(false)
	defer cleanup( )

	//! starting the gRPC server
	sharedUtils.CreateGRPCServer( )
}