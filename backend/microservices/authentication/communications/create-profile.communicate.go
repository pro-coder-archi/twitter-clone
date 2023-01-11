package communications

type CreateProfileEventPayload struct {

	Email string
	Password string
}

func CreateProfile(payload CreateProfileEventPayload) error {

	// TODO: implement

	return nil
}