package types

import "encoding/json"

// shape of object stored in redis, which represents temporary details of an user going through the registration process
type TemporaryUserDetailsRedisRecord struct {
	IsVerified bool

	Email string
	Name string
}

func(record TemporaryUserDetailsRedisRecord) MarshalBinary( ) ([]byte, error) {
	return json.Marshal(record)
}