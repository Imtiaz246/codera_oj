package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/imtiaz246/codera_oj/internal/adapters/handler/http/utils"
	"github.com/imtiaz246/codera_oj/internal/core/domain/dto"
	"github.com/o1egl/paseto"
	"net/http"
	"strings"
	"time"
)

// PasetoConfig indicates the configuration of the token middleware
type PasetoConfig struct {
	TokenLookup  string
	ContextKey   string
	Filter       func(ctx *fiber.Ctx) bool
	AuthScheme   string
	SymmetricKey []byte
}

// NewPasetoDefaultConfig returns default configuration for the token middleware
func newPasetoDefaultConfig(key []byte) *PasetoConfig {
	return &PasetoConfig{
		TokenLookup:  fmt.Sprintf("header:%s,cookie:token", fiber.HeaderAuthorization),
		ContextKey:   "payload",
		Filter:       nil,
		AuthScheme:   "Bearer",
		SymmetricKey: key,
	}
}

// NewPasetoMiddleware returns a paseto fiber middleware handler
func NewPasetoMiddleware(tokenConfig *dto.TokenConfig, configs ...*PasetoConfig) fiber.Handler {
	var c *PasetoConfig
	if len(configs) == 0 {
		c = newPasetoDefaultConfig(tokenConfig.Key)
	} else {
		c = configs[0]
	}
	tokenExtractor := createTokenExtractor(c.TokenLookup, c.AuthScheme)
	decryptor := paseto.NewV2()

	return func(ctx *fiber.Ctx) error {
		if c.Filter != nil && !c.Filter(ctx) {
			return ctx.Next()
		}
		token := tokenExtractor(ctx)
		if token == "" {
			return ctx.Status(http.StatusForbidden).JSON(utils.NewError(fmt.Errorf("token not found")))
		}

		pasetoPayload := new(paseto.JSONToken)
		err := decryptor.Decrypt(token, c.SymmetricKey, pasetoPayload, nil)
		if err != nil {
			return ctx.Status(http.StatusForbidden).JSON(utils.NewError(fmt.Errorf("invalid token")))
		}

		if time.Now().After(pasetoPayload.Expiration) {
			return ctx.Status(http.StatusForbidden).JSON(utils.NewError(fmt.Errorf("token has expired, expired time is: %v", pasetoPayload.Expiration)))
		}

		ctx.Locals(c.ContextKey, pasetoPayload)

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
				ss := strings.Split(ctx.Get(tokenSourceParts[1]), authScheme)
				if len(ss) > 1 {
					return ss[1]
				}
				return ""
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
