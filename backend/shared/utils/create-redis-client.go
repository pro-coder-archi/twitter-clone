package utils

import (
	"fmt"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

var (
	uri string
	password string

	miniRedisServer *miniredis.Miniredis
)

func CreateRedisClient(isForTesting bool) (*redis.Client, func( )) {
	const methodName= "CreateRedisClient"

	if !isForTesting {
		uri= GetEnv("REDIS_URL"); password= GetEnv("REDIS_PASSWORD")

	} else {
		var error error

		miniRedisServer, error= miniredis.Run( )
		if error != nil {
			Log(LogDetails{

				Method: methodName,
				Message: fmt.Sprintf("❌ error running mini-redis database server : \n%s", error.Error( )),
			})}

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
		Log(LogDetails{

			Method: methodName,
			Message: fmt.Sprintf("❌ error connecting to redis : \n%s", error.Error( )),
		})}

	Log(LogDetails{

		Type: DEFAULT_LOG,
		Method: methodName,
		Message: "🔥 successfully connected to redis database",
	})

	return redisClient,

		// cleanup function
		func( ) {

			defer miniRedisServer.Close( )
			defer redisClient.Close( )
		}
}