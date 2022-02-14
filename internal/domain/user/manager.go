package user

import "time"

type JWTManager interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
	ParseToken(accessToken string) (string, error)
}
