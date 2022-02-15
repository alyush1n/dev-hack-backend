package apperror

import "errors"

var (
	TokenMethodError = errors.New("unexpected token method")
	TokenError       = errors.New("fail to parse token")
	TokenClaimsError = errors.New("error get user claims from token")
	RandomizerError  = errors.New("error set refresh token")
)
