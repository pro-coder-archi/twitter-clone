package globals

import (
	"github.com/go-redis/redis"

	"authentication/repository"
)

type variables struct {

	Repository repository.Querier
	RedisClient *redis.Client
}

var Variables= &variables{ }