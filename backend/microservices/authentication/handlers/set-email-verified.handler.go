package handlers

import (
	"encoding/json"
	sharedUtils "shared/utils"

	"authentication/globals"
	proto "authentication/proto/generated"
	"authentication/types"
)

func SetEmailVerifiedHandler(setEmailVerifiedRequest *proto.SetEmailVerifiedRequest) (*proto.SetEmailVerifiedResponse, error) {

	var (
		methodName= "SetEmailVerifiedHandler"

		response *proto.SetEmailVerifiedResponse= nil
	)

	//! fetch the record from redis
	value, error := globals.Variables.RedisClient.Get(setEmailVerifiedRequest.Email).Result( )
	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Message: error,
			Method: methodName,
		})

		return response, error }

	//! unmarshalling and update the record

	var temporaryUserDetails types.TemporaryUserDetailsRedisRecord

	error= json.Unmarshal([]byte(value), &temporaryUserDetails)
	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Message: error,
			Method: methodName,
		})

		return response, error }

	temporaryUserDetails.IsVerified= true

	//! updating the record in redis
	error= globals.Variables.RedisClient.Set(setEmailVerifiedRequest.Email, temporaryUserDetails, -1).Err( )
	if error != nil {
		sharedUtils.Log(sharedUtils.LogDetails{

			Message: error,
			Method: methodName,
		})

		return response, error }

	return response, nil
}