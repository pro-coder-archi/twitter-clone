package middleware_tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"authentication/middlewares"
	"authentication/proto"
	"authentication/repository"
)

func TestStartRegistrationMiddleware(t *testing.T) {
	t.Parallel( )

	//! preparing the testcases

	type TestCase struct {

		description string
		input *proto.StartRegistrationRequest
		expectedOutput *string
		buildStub func( )
	}

	var testcases= []TestCase {
		{
			description: "🧪 invalid email should be detected",
			input: &proto.StartRegistrationRequest{
				Email: "invalid",
			},
			expectedOutput: &middlewares.InvalidEmailError,
		},
		{
			description: "🧪 invalid name should be detected",
			input: &proto.StartRegistrationRequest{
				Email: "archismanmridha12345@gmail.com",
				Name: "x",
			},
			expectedOutput: &middlewares.InvalidNameError,
		},
		{
			description: "🧪 duplicate email should be detected",
			input: &proto.StartRegistrationRequest{
				Email: "archismanmridha12345@gmail.com",
				Name: "archi",
			},
			expectedOutput: &middlewares.DuplicateEmailError,

			buildStub: func( ) {
				mockQuerier.
					EXPECT( ).
					FindRegisteredEmail(context.Background( ), "archismanmridha12345@gmail.com").
					Return(
						repository.AuthenticationUser{
							Email: "archismanmridha12345@gmail.com",
							Password: "password",

						}, nil,
					)
			},
		},
		{
			description: "🧪 request with valid input should pass through middleware successfully",
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

				output := middlewares.StartRegistrationMiddleware(testcase.input)

				if testcase.expectedOutput == nil {
					assert.Equal(t, testcase.expectedOutput, output) } else {

				assert.Equal(t, *output, *testcase.expectedOutput) }
			},
		)
	}
}