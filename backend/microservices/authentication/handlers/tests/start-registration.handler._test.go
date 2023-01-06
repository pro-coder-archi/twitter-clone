package handler_tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"authentication/global"
	"authentication/handlers"
	"authentication/middlewares"
	proto "authentication/proto/generated"
	"authentication/repository"
)

func TestStartRegistrationHandler(t *testing.T) {
	t.Parallel( )

	//! defining testcases

	type TestCase struct {

		description string
		input *proto.StartRegistrationRequest
		expectedOutput *proto.StartRegistrationResponse
		buildStub func( )
		testcaseContinuation func(input *proto.StartRegistrationRequest)
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
				Email: "archismanmridha12345@gmail.com",
				Name: "archi",
			},
			expectedOutput: nil,

			buildStub: func( ) {
				mockQuerier.
					EXPECT( ).
					FindRegisteredEmail(context.Background( ), "archismanmridha12345@gmail.com").
					Return(repository.AuthenticationUser{ }, sql.ErrNoRows)
			},

			testcaseContinuation: func(input *proto.StartRegistrationRequest) {
				value, error := global.GlobalVariables.RedisClient.Get(input.Email).Result( )

				assert.Nil(t, error)
				assert.NotEmpty(t, value)
			},
		},
	}

	//! running the testcases
	for _, item := range testcases {
		testcase := item

		t.Run(
			testcase.description, func(t *testing.T) {
				t.Parallel( )

				if testcase.buildStub != nil {
					testcase.buildStub( ) }

				output, error := handlers.StartRegistrationHandler(testcase.input)
				assert.Nil(t, error)

				if testcase.expectedOutput == nil {
					assert.Equal(t, testcase.expectedOutput, output) } else {

				assert.Equal(t, *output.Error, *testcase.expectedOutput.Error) }

				if testcase.testcaseContinuation != nil {
					testcase.testcaseContinuation(testcase.input) }
			},
		)
	}
}