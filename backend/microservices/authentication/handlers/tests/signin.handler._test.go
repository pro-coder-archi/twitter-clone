package handler_tests

import (
	"context"
	"database/sql"
	sharedUtils "shared/utils"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"

	"authentication/handlers"
	proto "authentication/proto/generated"
)

func TestSigninHandler(t *testing.T) {

	//! setup

	var randomEmail= faker.Email( )

	//! defining testcases

	type TestCase struct {

		description string
		input *proto.SigninRequest
		expectedOutput *proto.SigninResponse
		buildStub func( )
		continuation func(generatedJwt string)
	}

	testcases := []TestCase{
		{
			description: "🧪 user cannot signin without registration",
			input: &proto.SigninRequest{
				Email: randomEmail,
				Password: "password",
			},
			expectedOutput: &proto.SigninResponse{
				Error: &handlers.UserNotFoundError,
			},

			buildStub: func( ) {
				mockQuerier.
					EXPECT( ).
					GetPasswordForEmail(context.Background( ), randomEmail).
					Return("", sql.ErrNoRows)
			},
		},
		{
			description: "🧪 user cannot signin if password is wrong",
			input: &proto.SigninRequest{
				Email: randomEmail,
				Password: "password",
			},
			expectedOutput: &proto.SigninResponse{
				Error: &handlers.WrongPasswordError,
			},

			buildStub: func( ) {
				mockQuerier.
					EXPECT( ).
					GetPasswordForEmail(context.Background( ), randomEmail).
					Return("different password", nil)
			},
		},
		{
			description: "🧪 existing user with corret password should signin successfully",
			input: &proto.SigninRequest{
				Email: randomEmail,
				Password: "password",
			},

			buildStub: func( ) {
				mockQuerier.
					EXPECT( ).
					GetPasswordForEmail(context.Background( ), randomEmail).
					Return("password", nil)
			},

			continuation: func(jwt string) {
				isVerified, error := sharedUtils.VerifyJwt(jwt)

				assert.Nil(t, error)
				assert.True(t, isVerified)
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

				output, error := handlers.SigninHandler(testcase.input)
				assert.Nil(t, error)

				if item.expectedOutput != nil {
					assert.Equal(t, *testcase.expectedOutput.Error, *output.Error) } else {

				assert.Nil(t, output.Error) }

				if testcase.continuation != nil {
					testcase.continuation(output.Jwt) }
			},
		)
	}
}