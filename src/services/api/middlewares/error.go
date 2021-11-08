package middlewares

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/zikwall/monit/src/pkg/exceptions"
	"github.com/zikwall/monit/src/pkg/logger"
	"github.com/zikwall/monit/src/services/api/actions"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		code := fiber.StatusInternalServerError
		value := "Internal Server Error"

		var e *fiber.Error
		var w *exceptions.WrapError

		if errors.As(err, &e) {
			code = e.Code
			value = e.Message
		} else if errors.As(err, &w) {
			var pub *exceptions.ErrPublic
			var pri *exceptions.ErrPrivate

			if errors.As(err, &pub) {
				value = fmt.Sprintf("%s: %v", w.Context, pub.Error())
			} else if errors.As(err, &pri) {
				logger.Warning(fmt.Sprintf("%s: %v", w.Context, pri.Error()))
			}
		}

		return ctx.Status(code).JSON(actions.Response{
			"status":  code,
			"message": value,
		})
	}

	return nil
}
