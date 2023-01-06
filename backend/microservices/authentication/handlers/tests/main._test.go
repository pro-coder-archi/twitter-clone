package handler_tests

import (
	sharedUtils "shared/utils"
	"testing"

	"github.com/golang/mock/gomock"

	"authentication/global"
	"authentication/mocks"
)

var mockQuerier *mocks.MockQuerier

func TestMain(m *testing.M) {
	var cleanup func( )

	//! mocking redis database
	global.GlobalVariables.RedisClient, cleanup= sharedUtils.CreateRedisClient(true)
	defer cleanup( )

	//! mocking cockroach database

	mockingController := gomock.NewController(&testing.T{ })
	defer mockingController.Finish( )

	mockQuerier= mocks.NewMockQuerier(mockingController)

	global.GlobalVariables.Repository= mockQuerier

	//! running the tests
	m.Run( )
}