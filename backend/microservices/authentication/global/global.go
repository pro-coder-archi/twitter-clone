package global

import (
	"github.com/go-redis/redis"

	"authentication/repository"
)

var GlobalVariables= &globalVariables{ }

type globalVariables struct {

	Repository *repository.Queries
	RedisClient *redis.Client
}

var (
	ServerError= "server error occured"
)