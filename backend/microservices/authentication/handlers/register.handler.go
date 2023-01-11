package handlers

import (
	"context"
	"log"

	"authentication/communications"
	"authentication/global"
	proto "authentication/proto/generated"
	"authentication/repository"
)

func RegisterHandler(registerRequest *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	var response *proto.RegisterResponse= nil

	//! evicting the record from redis
	_, error := global.GlobalVariables.RedisClient.Del(registerRequest.Email).Result( )
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

		return response, error }

	//! sending request to profile service to create new profile

	communications.CreateProfile(
		communications.CreateProfileEventPayload{

			Email: registerRequest.Email,
			Password: registerRequest.Password,
		},
	)

	return response, nil
}