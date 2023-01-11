package configuration

type AppConfiguration struct {

	CockroachDBUri string `envconfig:"COCKROACHDB_URL" required:"true"`

	RedisUri string `envconfig:"REDIS_URL" required:"true"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" required:"true"`
}