package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mnabila/mytask/common"
	"github.com/mnabila/mytask/internal/entities"
)

// UseJWTMiddleware is a middleware function that validates and refreshes JWT tokens.
// It checks the "Authorization" header in the request, validates the JWT token,
// and refreshes the token if it is about to expire.
// The function takes a secret string parameter representing the JWT secret key.
// It returns a Fiber handler function that can be used as middleware.
func UseJWTMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the "Authorization" header from the request
		authorization := c.GetReqHeaders()[fiber.HeaderAuthorization]

		// Split the authorization header into fields
		authFields := strings.Fields(authorization)
		if len(authFields) < 2 {
			// Return unauthorized response if the authorization header is missing or incomplete
			return c.Status(fiber.StatusUnauthorized).JSON(
				entities.ApiResponse{
					Success: false,
					Message: jwt.ErrTokenSignatureInvalid.Error(),
				},
			)
		}

		if authFields[0] != "Bearer" {
			// Return unauthorized response if the authorization type is not "Bearer"
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: jwt.ErrTokenSignatureInvalid.Error(),
			})
		}

		// Unmarshal and validate the JWT claims using the provided secret
		claims, err := common.UnmarshalClaims(secret, authFields[1])
		if err != nil {
			// Return unauthorized response if the token is invalid or expired
			return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
				Success: false,
				Message: err.Error(),
			})
		}

		// Refresh the token if it is about to expire
		if time.Until(claims.ExpiresAt.Time) <= 5 {
			claims.ExpiresAt.Add(10 * time.Minute)
			authorization, err = common.MarshalClaims(secret, claims)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(entities.ApiResponse{
					Success: false,
					Message: err.Error(),
				})
			}
		}

		// Set the updated authorization header and store the claims in the context locals
		c.Set(fiber.HeaderAuthorization, authorization)
		c.Locals("claims", claims)

		return c.Next()
	}
}
