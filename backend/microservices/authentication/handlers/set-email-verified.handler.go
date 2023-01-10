package handlers

import (
	"encoding/json"
	"log"

	"authentication/global"
	proto "authentication/proto/generated"
	"authentication/types"
)

func SetEmailVerifiedHandler(
	setEmailVerifiedRequest *proto.SetEmailVerifiedRequest) (*proto.SetEmailVerifiedResponse, error) {

	//! fetch the record from redis
	value, error := global.GlobalVariables.RedisClient.Get(setEmailVerifiedRequest.Email).Result( )
	if error != nil {
		log.Println(error.Error( ))

		// TODO: handle error
		log.Println("⚠️ TODO: handle error when redis record not found for a user during the registration process")

		return nil, nil }

	//! unmarshalling and update the record

	var temporaryUserDetails types.TemporaryUserDetailsRedisRecord

	error= json.Unmarshal([]byte(value), &temporaryUserDetails)
	if error != nil {
		log.Println(error.Error( ))

		// TODO: handle error
		log.Println("⚠️ TODO: handle redis record unmarshalling error")

		return nil, nil }

	temporaryUserDetails.IsVerified= true

	//! updating the record in redis
	error= global.GlobalVariables.RedisClient.Set(setEmailVerifiedRequest.Email, temporaryUserDetails, -1).Err( )
	if error != nil {
		log.Println(error.Error( ))

		// TODO: handle error
		log.Println("⚠️ TODO: handle redis record update error")

		return nil, nil }

	return nil, nil
}