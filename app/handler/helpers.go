package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/imtiaz246/codera_oj/app/models"
	"github.com/imtiaz246/codera_oj/initializers/config"
	"time"
)

func ExtractRequestedUserFromClaims(c *fiber.Ctx) *requestedUser {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	email := claims["email"].(string)
	return &requestedUser{
		Username: username,
		Email:    email,
	}
}

func GenerateEncodedToken(u *models.User) (string, error) {
	claims := jwt.MapClaims{
		"email":        u.Email,
		"username":     u.Username,
		"organization": u.Organization,
		"firstName":    u.FirstName,
		"lastName":     u.LastName,
		"country":      u.Country,
		"exp":          time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, err := token.SignedString(config.GetAuthConfig().JWTSecret)
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}
