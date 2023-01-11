package handler_tests

import (
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"

	"authentication/handlers"
	proto "authentication/proto/generated"
	"authentication/repository"
)

func TestRegisterHandler(t *testing.T) {

	//! pre-requisities

	var randomEmail= faker.Email( )

	//! defining testcases

	type TestCase struct {

		description string
		input *proto.RegisterRequest
		expectedOutput *proto.RegisterResponse
		buildStub func( )
	}

	testcases := []TestCase {
		{
			description: "🧪 registration should be denied if time limit exceeded for the process",
			input: &proto.RegisterRequest{
				Email: randomEmail,
				Password: "password",
			},
			expectedOutput: &proto.RegisterResponse{ Error: &handlers.RegistrationTimeupError },
		},
		{
			description: "🧪 user should be registered successfully",
			input: &proto.RegisterRequest{
				Email: randomEmail,
				Password: "password",
			},
			expectedOutput: nil,

			buildStub: func( ) {
				CreateTemporaryUserDetailsRedisRecord(t, randomEmail)

				mockQuerier.
					EXPECT( ).
					CreateUser(context.Background( ), repository.CreateUserParams{

						Email: randomEmail,
						Password: "password",
					}).
					Return(nil)
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

				output, error := handlers.RegisterHandler(testcase.input)
				assert.Nil(t, error)

				assert.Equal(t, testcase.expectedOutput, output)
			},
		)
	}
}