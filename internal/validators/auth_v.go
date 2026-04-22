package validators

import (
	"errors"
	"mess-manager/internal/dto"
)

func ValidateLogIn(logINReq *dto.LogInReq)error {
	if  logINReq.UsernameOrEmail == "" {
		return errors.New("username or email is required")
	}
	return nil
}