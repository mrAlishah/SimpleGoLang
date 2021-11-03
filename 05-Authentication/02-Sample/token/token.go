package token

import (
	models "Authentication/models"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	jwtProviderToken = "SecretTokenSecretToken"
	ip               = "192.168.0.107"
)

func GenerateToken(claims *models.JwtClaims, expirationTime time.Time) (string, error) {

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString([]byte(jwtProviderToken))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
