package jwt

import (
	"dev-hack-backend/internal/service/user"
	"dev-hack-backend/pkg/apperror"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

type manager struct {
	signingKey string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewManager(signingKey string) user.JWTManager {
	return &manager{signingKey: signingKey}
}

func (m *manager) NewJWT(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *manager) NewRefreshToken() (string, error) {
	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", apperror.RandomizerError
	}

	return fmt.Sprintf("%x", bytes), nil
}

func (m *manager) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, apperror.TokenMethodError
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", apperror.TokenError
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", apperror.TokenClaimsError
	}

	return claims["sub"].(string), nil
}
