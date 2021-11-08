package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zikwall/monit/src/pkg/exceptions"
	"github.com/zikwall/monit/src/pkg/jwt"
)

func Access() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		claims, ok := ctx.Locals("claims").(*jwt.Claims)
		if !ok {
			return exceptions.Wrap("access", fiber.NewError(500, "Oops... Something went wrong.s"))
		}

		if claims.UUID == 0 {
			return exceptions.Wrap("access", fiber.NewError(403, "It seems you don't have access."))
		}
		return ctx.Next()
	}
}
