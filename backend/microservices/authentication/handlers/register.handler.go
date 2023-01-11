package handlers

import (
	"context"
	"encoding/json"
	"log"
	sharedErrors "shared/errors"

	"authentication/communications"
	"authentication/global"
	proto "authentication/proto/generated"
	"authentication/repository"
	"authentication/types"
)

func RegisterHandler(registerRequest *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	var response *proto.RegisterResponse= nil

	value, error := global.GlobalVariables.RedisClient.Get(registerRequest.Email).Result( )
	if error != nil {
		log.Println(error.Error( ))

		return &proto.RegisterResponse{ Error: &RegistrationTimeupError }, nil }

	var temporaryUserDetails types.TemporaryUserDetailsRedisRecord

	error= json.Unmarshal([]byte(value), &temporaryUserDetails)
	if error != nil {
		log.Println(error.Error( ))

		return &proto.RegisterResponse{ Error: &sharedErrors.ServerError }, nil }

	//! evicting the record from redis
	_, error= global.GlobalVariables.RedisClient.Del(registerRequest.Email).Result( )
	if error != nil {
		log.Println(error.Error( )) }

	//! saving the user details permanently in cockroachDB

	error= global.GlobalVariables.Repository.CreateUser(
		context.Background( ), repository.CreateUserParams{

			Email: registerRequest.Email,
			Password: registerRequest.Password,
		},
	)

	if error != nil {
		log.Println(error.Error( ))

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