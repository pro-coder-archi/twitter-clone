package utils

import (
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

var (
	uri string
	password string

	miniRedisServer *miniredis.Miniredis
)

func CreateRedisClient(isForTesting bool) (*redis.Client, func( )) {

	if !isForTesting {
		uri= GetEnv("REDIS_URL"); password= GetEnv("REDIS_PASSWORD")

	} else {
		var error error

		miniRedisServer, error= miniredis.Run( )
		if error != nil {
			log.Println("❌ error running mini-redis database server")

			log.Println(error.Error( )) }

		uri= miniRedisServer.Addr( )
	}

	redisClient := redis.NewClient(
		&redis.Options{

			Addr: uri,
			Password: password,
		},
	)

	_, error := redisClient.Ping( ).Result( )
	if(error != nil) {
		log.Fatalf("❌ error connecting to redis : ")

		log.Fatalf(error.Error( )) }

	log.Println("🔥 successfully connected to redis database")

	return redisClient,

		// cleanup function
		func( ) {

			defer miniRedisServer.Close( )
			defer redisClient.Close( )
		}
}