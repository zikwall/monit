package middlewares

import (
	"github.com/gofiber/fiber/v2"

	"github.com/zikwall/monit/src/pkg/jwt"
)

func JWT(public []byte) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var claims *jwt.Claims

		if token, ok := jwt.AuthHeader(ctx.Get(jwt.AuthHeaderName)); ok {
			cl, err := jwt.VerifyJwtToken(token, public)
			if err == nil {
				claims = cl
			}
		}

		if claims == nil {
			claims = &jwt.Claims{}
		}

		ctx.Locals("claims", claims)
		return ctx.Next()
	}
}
