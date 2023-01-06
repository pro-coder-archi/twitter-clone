package middleware_tests

import (
	"testing"

	"github.com/golang/mock/gomock"

	"authentication/global"
	"authentication/mocks"
)

var mockQuerier *mocks.MockQuerier

func TestMain(m *testing.M) {

	//* mocking database call

	mockingController := gomock.NewController(&testing.T{ })
	defer mockingController.Finish( )

	mockQuerier= mocks.NewMockQuerier(mockingController)

	global.GlobalVariables.Repository= mockQuerier

	//* running the tests
	m.Run( )
}