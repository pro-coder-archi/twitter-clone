package handlers

import (
	"context"
	"encoding/json"
	"log"
	"shared/communications"

	"authentication/global"
	proto "authentication/proto/generated"
	"authentication/repository"
	"authentication/types"
)

func RegisterHandler(registerRequest *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	var response *proto.RegisterResponse= nil

	//! fetch the record from redis
	value, error := global.GlobalVariables.RedisClient.Get(registerRequest.Email).Result( )
	if error != nil {
		log.Println(error.Error( ))

		communications.ReportError(error)
		return response, nil
	}

	//! unmarshalling the record

	var temporaryUserDetails types.TemporaryUserDetailsRedisRecord

	error= json.Unmarshal([]byte(value), &temporaryUserDetails)
	if error != nil {
		log.Println(error.Error( ))

		communications.ReportError(error)
		return response, nil
	}

	//! saving the user details permanently in cockroachDB

	error= global.GlobalVariables.Repository.CreateUser(
		context.Background( ), repository.CreateUserParams{

			Email: temporaryUserDetails.Email,
			Password: registerRequest.Password,
		},
	)

	if error != nil {
		log.Println(error.Error( ))

		communications.ReportError(error)
		return response, nil
	}

	return response, nil
}