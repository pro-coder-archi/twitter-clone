package communications

type CreateProfileEventPayload struct {

	Name string
	Email string
	Password string
}

func CreateProfile(payload CreateProfileEventPayload) error {

	// TODO: implement

	return nil
}