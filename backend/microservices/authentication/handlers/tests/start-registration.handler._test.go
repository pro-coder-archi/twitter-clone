package handler_tests

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"

	"authentication/global"
	"authentication/handlers"
	"authentication/middlewares"
	proto "authentication/proto/generated"
	"authentication/repository"
	"authentication/types"
)

func TestStartRegistrationHandler(t *testing.T) {

	//! setup

	var randomEmail= faker.Email( )

	//! defining testcases

	type TestCase struct {

		description string
		input *proto.StartRegistrationRequest
		expectedOutput *proto.StartRegistrationResponse
		buildStub func( )
		continuation func(input *proto.StartRegistrationRequest)
	}

	testcases := []TestCase{
		{
			description: "🧪 invalid request should not pass through the middleware",
			input: &proto.StartRegistrationRequest{
				Email: "invalid email",
				Name: "archi",
			},
			expectedOutput: &proto.StartRegistrationResponse{ Error: &middlewares.InvalidEmailError },
		},
		{
			description: "🧪 valid request should pass through the middleware and details should be saved in redis",
			input: &proto.StartRegistrationRequest{
				Email: randomEmail,
				Name: "archi",
			},
			expectedOutput: nil,

			buildStub: func( ) {
				mockQuerier.
					EXPECT( ).
					FindRegisteredEmail(context.Background( ), randomEmail).
					Return(repository.User{ }, sql.ErrNoRows)
			},

			continuation: func(input *proto.StartRegistrationRequest) {
				value, error := global.GlobalVariables.RedisClient.Get(input.Email).Result( )

				assert.Nil(t, error)
				assert.NotEmpty(t, value)

				var temporaryUserDetails types.TemporaryUserDetailsRedisRecord

				error= json.Unmarshal([]byte(value), &temporaryUserDetails)
				assert.Nil(t, error)

				assert.False(t, temporaryUserDetails.IsVerified)
				assert.Equal(t, input.Email, temporaryUserDetails.Email)
				assert.Equal(t, input.Name, temporaryUserDetails.Name)
			},
		},
	}

	//! running the testcases
	for _, item := range testcases {
		testcase := item

		t.Run(
			testcase.description, func(t *testing.T) {

				if testcase.buildStub != nil {
					testcase.buildStub( ) }

				output, error := handlers.StartRegistrationHandler(testcase.input)
				assert.Nil(t, error)

				if testcase.expectedOutput == nil {
					assert.Equal(t, testcase.expectedOutput, output) } else {

				assert.Equal(t, *output.Error, *testcase.expectedOutput.Error) }

				if testcase.continuation != nil {
					testcase.continuation(testcase.input) }
			},
		)
	}
}