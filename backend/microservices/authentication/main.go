package main

import (
	"context"
	"log"
	sharedUtils "shared/utils"

	"github.com/go-redis/redis"

	"authentication/global"
	"authentication/repository"
)

func main( ) {

	ctx, cancel := context.WithCancel(context.Background( ))
	defer cancel( )

	//* connecting to cockroachDB
	dbConnection := sharedUtils.CreateCockroachDBConnection(ctx)
	global.GlobalVariables.Repository = repository.New(dbConnection)

	//* connecting to redis

	redisClient := redis.NewClient(
		&redis.Options{
			Addr: sharedUtils.GetEnv("REDIS_URL"),
			Password: sharedUtils.GetEnv("REDIS_PASSWORD"),
		},
	)

	_, error := redisClient.Ping( ).Result( )

	if(error != nil) {
		log.Fatalf("❌ error connecting to redis : ")

		log.Fatalf(error.Error( )) }

	log.Println("🔥 successfully connected to redis database")

	global.GlobalVariables.RedisClient= redisClient

	//* starting the gRPC server
	sharedUtils.CreateGRPCServer( )
}