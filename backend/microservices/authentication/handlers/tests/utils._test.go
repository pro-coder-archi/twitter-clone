package handler_tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"authentication/globals"
	"authentication/types"
)

func CreateTemporaryUserDetailsRedisRecord(t *testing.T, email string) {

	temporaryUserDetails, error := json.Marshal(
		types.TemporaryUserDetailsRedisRecord {
			IsVerified: true,

			Email: email,
			Name: "archi",
		},
	)
	assert.Nil(t, error)

	error= globals.Variables.RedisClient.Set(email, temporaryUserDetails, 300 * time.Second).Err( )
	assert.Nil(t, error)
}