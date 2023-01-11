package handler_tests

import (
	sharedUtils "shared/utils"
	"testing"

	"github.com/golang/mock/gomock"

	"authentication/globals"
	"authentication/mocks"
)

var mockQuerier *mocks.MockQuerier

func TestMain(m *testing.M) {
	var cleanup func( )

	//! mocking redis database
	globals.Variables.RedisClient, cleanup= sharedUtils.CreateRedisClient(true)
	defer cleanup( )

	//! mocking cockroach database

	mockingController := gomock.NewController(&testing.T{ })
	defer mockingController.Finish( )

	mockQuerier= mocks.NewMockQuerier(mockingController)

	globals.Variables.Repository= mockQuerier

	//! running the tests
	m.Run( )
}