package forms

import (
	"errors"

	"github.com/zikwall/monit/src/pkg/exceptions"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *Login) Validate() error {
	if l.Username != "" && l.Password != "" {
		return nil
	}

	return exceptions.ThrowPublicError(errors.New("username or password is empty"))
}
