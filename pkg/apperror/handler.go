package apperror

import "errors"

var (
	HeaderEmptyError   = errors.New("header is empty")
	InvalidHeaderError = errors.New("invalid auth header")
	TokenEmptyError    = errors.New("token is empty")
	BadData            = errors.New("not all parameters are specified: ")
)
