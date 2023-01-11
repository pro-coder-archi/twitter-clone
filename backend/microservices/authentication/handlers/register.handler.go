package handlers

import (
	"context"
	"encoding/json"
	sharedErrors "shared/errors"
	sharedUtils "shared/utils"

	"authentication/communications"
	"authentication/global"
	proto "authentication/proto/generated"
	"authentication/repository"
	"authentication/types"
)

func RegisterHandler(registerRequest *proto.RegisterRequest) (*proto.RegisterResponse, error) {

	var (
		methodName= "RegisterHandler"

		response *proto.RegisterResponse= nil
	)

	//! fetching temporary user details from redis

	value, error := global.GlobalVariables.RedisClient.Get(registerRequest.Email).Result( )
	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Message: error,
			Method: methodName,
		})

		return &proto.RegisterResponse{ Error: &RegistrationTimeupError }, nil }

	var temporaryUserDetails types.TemporaryUserDetailsRedisRecord

	error= json.Unmarshal([]byte(value), &temporaryUserDetails)
	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Message: error,
			Method: methodName,
		})

		return &proto.RegisterResponse{ Error: &sharedErrors.ServerError }, nil }

	//! evicting the record from redis
	_, error= global.GlobalVariables.RedisClient.Del(registerRequest.Email).Result( )
	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Message: error,
			Method: methodName,
		})}

	//! saving the user details permanently in cockroachDB

	error= global.GlobalVariables.Repository.CreateUser(
		context.Background( ), repository.CreateUserParams{

			Email: registerRequest.Email,
			Password: registerRequest.Password,
		},
	)

	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Message: error,
			Method: methodName,
		})

		return &proto.RegisterResponse{ Error: &sharedErrors.ServerError }, nil }

	//! sending request to profile service to create new profile

	communications.CreateProfile(
		communications.CreateProfileEventPayload{

			Name: temporaryUserDetails.Name,
			Email: registerRequest.Email,
			Password: registerRequest.Password,
		},
	)

	return response, nil
}

var (
	RegistrationTimeupError= "please restart the registration process"
)