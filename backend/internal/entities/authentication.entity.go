package entities

import "github.com/golang-jwt/jwt/v5"

type AuthenticationRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UserClaims struct {
	Id string `json:"id,omitempty"`
	jwt.RegisteredClaims
}
