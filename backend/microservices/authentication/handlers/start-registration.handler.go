package handlers

import "authentication/proto"

func StartRegistrationHandler(
	startRegistrationRequest *proto.StartRegistrationRequest) (*proto.StartRegistrationResponse, error) {

	// check if email is not pre-registered using the authentication database

	return &proto.StartRegistrationResponse{ }, nil
}