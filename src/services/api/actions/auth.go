package actions

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/zikwall/monit/src/pkg/exceptions"
	"github.com/zikwall/monit/src/pkg/jwt"
	"github.com/zikwall/monit/src/services/api/constants"
	"github.com/zikwall/monit/src/services/api/forms"
)

func (ht *HTTPController) Login(ctx *fiber.Ctx) error {
	form := &forms.Login{}

	if err := ctx.BodyParser(&form); err != nil {
		return exceptions.Wrap("failed parse form body", err)
	}

	if err := form.Validate(); err != nil {
		return exceptions.Wrap("failed validate form", err)
	}

	if form.Username != constants.UserName || form.Password != constants.UserPass {
		return exceptions.Wrap("failed find user", exceptions.ThrowPublicError(
			errors.New("user not exist"),
		))
	}

	token, err := jwt.CreateJwtToken(&jwt.Claims{UUID: constants.UUID}, 1000, constants.PrivateKey)
	if err != nil {
		return exceptions.Wrap("invalid token", err)
	}

	return ctx.JSON(Response{
		"token": token,
	})
}
