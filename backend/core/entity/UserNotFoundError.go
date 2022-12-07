package entity

type UserNotFoundError struct {
	inner
	userID string
}

func (err UserNotFoundError) New(userID string) *UserNotFoundError {
	err.inner = innerFromString("user '%s' was not found", userID)
	err.userID = userID
	return &err
}

func (err *UserNotFoundError) UserID() string {
	return err.userID
}
