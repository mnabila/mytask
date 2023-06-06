package common

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mnabila/mytask/internal/entities"
)

// UnmarshalClaims parses a JWT token string and populates the provided claims struct with the extracted claims.
// It takes two parameters: `secret` (string) representing the JWT secret key used for token verification,
// and `tokenString` (string) representing the JWT token string to be parsed.
// The function returns a pointer to an `entities.UserClaims` struct containing the extracted claims or an error if the token parsing fails.
func UnmarshalClaims(secret, tokenString string) (*entities.UserClaims, error) {
	// Create a new empty claims struct to be populated
	claims := new(entities.UserClaims)

	// Parse the JWT token with the provided claims struct and secret key
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Return the secret key as the verification key
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, jwt.ErrTokenNotValidYet
	}

	// Return the populated claims struct
	return claims, nil
}

// MarshalClaims creates a JWT token string by marshaling the provided claims into a token and signing it with the given secret key.
// It takes a secret string parameter representing the JWT secret key used for token signing, and a pointer to an entities.UserClaims struct containing the claims.
// The function returns the JWT token string prefixed with "Bearer" or an error if the token signing fails.
func MarshalClaims(secret string, claims *entities.UserClaims) (string, error) {
	// Create a new token with the provided signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// Return the formatted token string prefixed with "Bearer"
	return fmt.Sprintf("Bearer %s", tokenString), nil
}
