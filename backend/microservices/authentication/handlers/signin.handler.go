package handlers

import (
	"context"
	"database/sql"
	sharedErrors "shared/errors"
	sharedUtils "shared/utils"

	"authentication/global"
	proto "authentication/proto/generated"
)

func SigninHandler(registerRequest *proto.SigninRequest) (*proto.SigninResponse, error) {

	//! search user from authentication database

	password, error := global.GlobalVariables.Repository.GetPasswordForEmail(context.Background( ), registerRequest.Email)

	if error == sql.ErrNoRows {
		return &proto.SigninResponse{ Error: &UserNotFoundError }, nil }

	if error != nil {
		return &proto.SigninResponse{ Error: &sharedErrors.ServerError }, nil }

	if password != registerRequest.Password {
		return &proto.SigninResponse{ Error: &WrongPasswordError }, nil }

	//! create jwt
	jwt, createJwtError := sharedUtils.CreateJwt(registerRequest.Email)

	response := proto.SigninResponse{

		Jwt: jwt,
		Error: createJwtError,
	}

	return &response, nil
}

var (
	UserNotFoundError= "user not found"
	WrongPasswordError= "wrong password provided"
)