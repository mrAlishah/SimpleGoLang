package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// JWTManager is a JSON web token manager
type JWTManager struct {
	//The secret key to sign and verify the access token,
	secretKey string
	//the valid duration of the token.
	tokenDuration time.Duration
}

// UserClaims is a custom JWT claims that contains some user's information
type UserClaims struct {
	//standard claims has several useful information that we can set
	jwt.StandardClaims
	//The JSON web token should contain a claims object, which has some useful information about the user who owns it
	Username string `json:"username"`
	Role     string `json:"role"`
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

// Generate generates and signs a new token for a user
func (manager *JWTManager) Generate(user *User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	//generate a token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//sign the generated token with your secret key. This will make sure that no one can fake an access token, since they don’t have your secret key.
	return token.SignedString([]byte(manager.secretKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {

	//We just have to call jwt.ParseWithClaims(), pass in the access token, an empty user claims, and a custom key function.
	//In this function, It’s very important to check the signing method of the token to make sure that it matches with the algorithm our server uses,
	//which in our case is HMAC. If it matches, then we just return the secret key that is used to sign the token.
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	//we get the claims from the token and convert it to a UserClaims
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
