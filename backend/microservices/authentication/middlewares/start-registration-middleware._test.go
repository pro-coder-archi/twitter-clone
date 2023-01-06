package middlewares

import (
	"context"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"authentication/global"
	"authentication/mocks"
	"authentication/proto"
	"authentication/repository"
)

func TestStartRegistrationMiddleware(t *testing.T) {

	//! mocking database call

	mockingController := gomock.NewController(t)
	defer mockingController.Finish( )

	var mockQuerier *mocks.MockQuerier= mocks.NewMockQuerier(mockingController)

	global.GlobalVariables.Repository= mockQuerier

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
			expectedOutput: &invalidEmailError,
		},
		{
			description: "🧪 invalid name should be detected",
			input: &proto.StartRegistrationRequest{
				Email: "archismanmridha12345@gmail.com",
				Name: "x",
			},
			expectedOutput: &invalidNameError,
		},
		{
			description: "🧪 duplicate email should be detected",
			input: &proto.StartRegistrationRequest{
				Email: "archismanmridha12345@gmail.com",
				Name: "archi",
			},
			expectedOutput: &duplicateEmailError,

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
	for _, testcase := range testcases {
		t.Run(
			testcase.description, func(t *testing.T) {

				if testcase.buildStub != nil {
					testcase.buildStub( ) }

				output := StartRegistrationMiddleware(testcase.input)

				if testcase.expectedOutput == nil {
					assert.Equal(t, testcase.expectedOutput, output) } else {

				assert.Equal(t, *output, *testcase.expectedOutput) }
			},
		)
	}
}