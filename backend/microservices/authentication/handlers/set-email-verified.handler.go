package handlers

import (
	"encoding/json"
	"log"
	"shared/communications"

	"authentication/global"
	proto "authentication/proto/generated"
	"authentication/types"
)

func SetEmailVerifiedHandler(setEmailVerifiedRequest *proto.SetEmailVerifiedRequest) (*proto.SetEmailVerifiedResponse, error) {
	var response *proto.SetEmailVerifiedResponse= nil

	//! fetch the record from redis
	value, error := global.GlobalVariables.RedisClient.Get(setEmailVerifiedRequest.Email).Result( )
	if error != nil {
		log.Println(error.Error( ))

		communications.ReportError(error)
		return response, nil }

	//! unmarshalling and update the record

	var temporaryUserDetails types.TemporaryUserDetailsRedisRecord

	error= json.Unmarshal([]byte(value), &temporaryUserDetails)
	if error != nil {
		log.Println(error.Error( ))

		communications.ReportError(error)
		return response, nil }

	temporaryUserDetails.IsVerified= true

	//! updating the record in redis
	error= global.GlobalVariables.RedisClient.Set(setEmailVerifiedRequest.Email, temporaryUserDetails, -1).Err( )
	if error != nil {
		log.Println(error.Error( ))

		communications.ReportError(error)
		return response, nil }

	return response, nil
}