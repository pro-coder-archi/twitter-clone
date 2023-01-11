package middleware_tests

import (
	"testing"

	"github.com/golang/mock/gomock"

	"authentication/globals"
	"authentication/mocks"
)

var mockQuerier *mocks.MockQuerier

func TestMain(m *testing.M) {

	//* mocking cockroach database

	mockingController := gomock.NewController(&testing.T{ })
	defer mockingController.Finish( )

	mockQuerier= mocks.NewMockQuerier(mockingController)

	globals.Variables.Repository= mockQuerier

	//* running the tests
	m.Run( )
}