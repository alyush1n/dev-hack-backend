package jwt

import (
	"dev-hack-backend/internal/domain/user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

const (
	tokenMethodError = "unexpected token method: %s"
	tokenError       = "fail to parse token with error %w"
	tokenClaimsError = "error get user claims from token"
	randomizerError  = "error set refresh token"
)

type manager struct {
	signingKey string
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

	source := rand.NewSource(time.Now().Unix())
	randomizer := rand.New(source)

	_, err := randomizer.Read(bytes)
	if err != nil {
		return "", fmt.Errorf(randomizerError)
	}

	return fmt.Sprintf("%x", bytes), nil
}

func (m *manager) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf(tokenMethodError, t.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", fmt.Errorf(tokenError, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf(tokenClaimsError)
	}

	return claims["sub"].(string), nil
}
