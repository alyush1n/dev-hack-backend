package apperror

import "errors"

var (
	MongoInsertError        = errors.New("failed to insert user")
	MongoUpdateError        = errors.New("failed to update user")
	MongoDeleteError        = errors.New("failed to delete user")
	MongoFindError          = errors.New("failed to find user")
	MongoUpdateSessionError = errors.New("failed to update session user")
	ObjectIdError           = errors.New("error with convert string to ObjectId")
	RefreshEqualError       = errors.New("refresh token is invalid")
	RefreshTimeError        = errors.New("failed to refresh token (time is up)")
	ClientError             = errors.New("failed to create client to mongodb")
)
