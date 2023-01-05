package middlewares

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"authentication/proto"
)

func TestStartRegistrationMiddleware(t *testing.T) {

	type TestCase struct {

		description string
		input *proto.StartRegistrationRequest
		expectedOutput string
	}

	var testcases= []TestCase {
		{
			description: "🧪 invalid email should be detected",
			input: &proto.StartRegistrationRequest{
				Email: "invalid",
			},
			expectedOutput: invalidEmailError,
		},
		{
			description: "🧪 invalid name should be detected",
			input: &proto.StartRegistrationRequest{
				Email: "archismanmridha12345@gmail.com",
				Name: "x",
			},
			expectedOutput: invalidNameError,
		},
		{
			description: "🧪 duplicate email should be detected",
			input: &proto.StartRegistrationRequest{
				Email: "archismanmridha12345@gmail.com",
				Name: "archi",
			},
			expectedOutput: duplicateEmailError,
		},
	}

	// TODO: mocking database call

	//* running the testcases
	for _, testcase := range testcases {
		t.Run(
			testcase.description, func(t *testing.T) {

				output := StartRegistrationMiddleware(testcase.input)
				assert.Equal(t, *output, testcase.expectedOutput)
			},
		)
	}
}