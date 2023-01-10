package handler_tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"

	"authentication/global"
	"authentication/handlers"
	proto "authentication/proto/generated"
	"authentication/types"
)

func TestSetEmailVerifiedHandler(t *testing.T) {
	t.Parallel( )

	//! pre-requisities

	var (
		randomEmail= faker.Email( )
	)

	temporaryUserDetails, error := json.Marshal(
		types.TemporaryUserDetailsRedisRecord {
			IsVerified: false,

			Email: randomEmail,
			Name: "archi",
		},
	)
	assert.Nil(t, error)

	error= global.GlobalVariables.RedisClient.Set(randomEmail, temporaryUserDetails, 300 * time.Second).Err( )
	assert.Nil(t, error)

	//! defining testcases

	type TestCase struct {

		description string
		input *proto.SetEmailVerifiedRequest
		expectedOutput *proto.SetEmailVerifiedResponse
		continuation func(input *proto.SetEmailVerifiedRequest)
	}

	testcases := []TestCase {
		{
			description: "🧪 user should be set as verified in redis record",
			input: &proto.SetEmailVerifiedRequest{
				Email: randomEmail,
			},
			expectedOutput: nil,
			continuation: func(input *proto.SetEmailVerifiedRequest) {
				value, error := global.GlobalVariables.RedisClient.Get(input.Email).Result( )

				assert.Nil(t, error)
				assert.NotEmpty(t, value)

				var temporaryUserDetails types.TemporaryUserDetailsRedisRecord

				error= json.Unmarshal([]byte(value), &temporaryUserDetails)
				assert.Nil(t, error)

				assert.True(t, temporaryUserDetails.IsVerified)
				assert.Equal(t, input.Email, temporaryUserDetails.Email)
			},
		},
	}

	//! running the testcases
	for _, item := range testcases {
		testcase := item

		t.Run(
			testcase.description, func(t *testing.T) {
				t.Parallel( )

				output, error := handlers.SetEmailVerifiedHandler(testcase.input)
				assert.Nil(t, error)

				assert.Equal(t, testcase.expectedOutput, output)

				if testcase.continuation != nil {
					testcase.continuation(testcase.input) }
			},
		)
	}
}