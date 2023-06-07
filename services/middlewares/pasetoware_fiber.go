package middlewares

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/initializers/config"
	"github.com/o1egl/paseto"
	"strings"
	"time"
)

// Config indicates the configuration of the token middleware
type Config struct {
	TokenLookup  string
	ContextKey   string
	Filter       func(ctx *fiber.Ctx) bool
	AuthScheme   string
	SymmetricKey string
}

// NewPasetoDefaultConfig returns default configuration for the token middleware
func NewPasetoDefaultConfig() *Config {
	authConfig := config.GetAuthConfig()
	return &Config{
		TokenLookup:  fmt.Sprintf("header:%s,cookie:token", fiber.HeaderAuthorization),
		ContextKey:   "payload",
		Filter:       nil,
		AuthScheme:   "Bearer",
		SymmetricKey: authConfig.Key,
	}
}

// New returns the fiber middleware handler
func New(config *Config) fiber.Handler {
	tokenExtractor := createTokenExtractor(config.TokenLookup, config.AuthScheme)
	decryptor := paseto.NewV2()

	return func(ctx *fiber.Ctx) error {
		if config.Filter != nil && !config.Filter(ctx) {
			return ctx.Next()
		}
		token := tokenExtractor(ctx)
		if token == "" {
			return fmt.Errorf("token not found")
		}

		pasetoPayload := new(paseto.JSONToken)
		key, err := base64.StdEncoding.DecodeString(config.SymmetricKey)
		if err != nil {
			return fmt.Errorf("token decoding failed: %v", err)
		}

		err = decryptor.Decrypt(token, key, pasetoPayload, nil)
		if err != nil {
			return fmt.Errorf("invalid token: %v", err)
		}

		if time.Now().After(pasetoPayload.Expiration) {
			return fmt.Errorf("token has expired, expired time is: %v", pasetoPayload.Expiration)
		}

		ctx.Locals(config.ContextKey, pasetoPayload)

		return ctx.Next()
	}
}

// createTokenExtractor creates a fiber handler to extract token from various source.
// The source could be anything where a token can be existed.
func createTokenExtractor(tokenLookup, authScheme string) func(ctx *fiber.Ctx) string {
	tokenSources := strings.Split(tokenLookup, ",")
	checks := make([]func(ctx *fiber.Ctx) string, 0)
	authScheme = strings.Split(authScheme, " ")[0] + " "

	for _, tokenSource := range tokenSources {
		tokenSourceParts := strings.Split(tokenSource, ":")
		switch tokenSourceParts[0] {
		case "header":
			checks = append(checks, func(ctx *fiber.Ctx) string {
				token := strings.Split(ctx.Get(tokenSourceParts[1]), authScheme)[1]
				return token
			})
		case "query":
			checks = append(checks, func(ctx *fiber.Ctx) string {
				return ctx.Query(tokenSourceParts[1])
			})
		case "params":
			checks = append(checks, func(ctx *fiber.Ctx) string {
				return ctx.Params(tokenSourceParts[1])
			})
		case "cookie":
			checks = append(checks, func(ctx *fiber.Ctx) string {
				return ctx.Cookies(tokenSourceParts[1])
			})
		}
	}

	return func(ctx *fiber.Ctx) string {
		for _, check := range checks {
			token := check(ctx)
			if token != "" {
				return token
			}
		}
		return ""
	}
}
