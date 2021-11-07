package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/services/api/actions"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		code := fiber.StatusInternalServerError
		value := "Internal Server Error"

		var e *fiber.Error

		if errors.As(err, &e) {
			code = e.Code
			value = e.Message
		}

		logger.Warning(err.Error())

		return ctx.Status(code).JSON(actions.Response{
			"status":  false,
			"message": value,
		})
	}

	return nil
}
