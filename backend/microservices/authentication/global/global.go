package global

import (
	"github.com/go-redis/redis"

	"authentication/repository"
)

type globalVariables struct {

	Repository repository.Querier
	RedisClient *redis.Client
}

var GlobalVariables= &globalVariables{ }

var (
	ServerError= "server error occured"
)